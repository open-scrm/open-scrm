package service

import (
	"context"
	"fmt"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/redikeys"
	"github.com/open-scrm/open-scrm/lib/redislock"
	"github.com/open-scrm/open-scrm/lib/utils"
	"github.com/open-scrm/open-scrm/lib/wxwork"
	"github.com/open-scrm/open-scrm/pkg/addressbook/dao"
	admodel "github.com/open-scrm/open-scrm/pkg/addressbook/model"
	dao2 "github.com/open-scrm/open-scrm/pkg/externalcontact/dao"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/model"
	"github.com/open-scrm/open-scrm/pkg/system/service"
	"github.com/pkg/errors"
)

type SyncExternalContactService struct {
	lock *redislock.RedisLock
}

func NewSyncExternalContactService() *SyncExternalContactService {
	return &SyncExternalContactService{
		lock: redislock.NewRedisLock(global.GetRedis()),
	}
}

func (s *SyncExternalContactService) SyncAll(ctx context.Context) error {
	ok, err := s.lock.GetLock(ctx, redikeys.SyncExternalContactKey, redikeys.SyncExternalContactExpire)
	if err != nil {
		return errors.Wrap(err, "获取全局锁失败")
	}
	if !ok {
		return fmt.Errorf("上一次同步未结束")
	}
	defer s.lock.UnLock(ctx, redikeys.SyncExternalContactKey)

	// sync All. 同步处理. 同步员工和客户的关系.
	var (
		lastId       int64
		syncUserSize int64 = 100
	)
	userDao := dao.NewUserDao(ctx)
	for {
		users, err := userDao.ScrollList(ctx, lastId, syncUserSize)
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("查询员工信息失败")
			return errors.Wrap(err, "查询员工信息失败")
		}
		if len(users) == 0 {
			break
		}
		if err := s.SyncByUserIds(ctx, users); err != nil {
			log.WithContext(ctx).WithError(err).WithField("users", users).Errorf("同步客户信息失败")
			return err
		}
		lastId = users[len(users)-1].Id
	}
	return nil
}

func (s *SyncExternalContactService) SyncByUserIds(ctx context.Context, users []*admodel.User) error {
	var (
		cursor        string
		limit         int64 = 100
		customerDao         = dao2.NewCustomerDao()
		followUserDao       = dao2.NewCustomerFollowUserDao()
		counter       int64
		usersMap      = map[string]*admodel.User{}
		userIds       []string
	)
	talentInfo, err := service.NewTalentService().GetTalentInfo(ctx)
	if err != nil {
		return errors.Wrap(err, "获取租户信息失败")
	}
	cli := global.GetWxWorkClient()

	token, err := cli.GetAddressBookToken(ctx, talentInfo.CorpId, talentInfo.AddressBookSecret)
	if err != nil {
		return errors.Wrap(err, "获取access token失败")
	}
	ctx = wxwork.ContextWithAccessToken(ctx, token)

	for _, user := range users {
		userIds = append(userIds, user.Userid)
		usersMap[user.Userid] = user
	}

	log.WithContext(ctx).WithField("userids", userIds).Infof("开始同步员工客户")
	for {
		resp, err := cli.BatchGetExternalContactList(ctx, userIds, cursor, limit)
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("批量获取联系人详情失败: %v %v %v", userIds, cursor, limit)
			return err
		}
		cursor = resp.NextCursor
		var (
			customers []*model.Customer
			followers []*model.CustomerFollowUser
		)
		for _, list := range resp.ExternalContactList {
			var (
				owner int64
			)
			sysUser, ok := usersMap[list.FollowInfo.Userid]
			if ok {
				owner = sysUser.Id
			}
			cus := &model.Customer{
				Id:             int64(global.GetSnowflakeNode().Generate()),
				ExternalUserId: list.ExternalContact.ExternalUserid,
				Name:           list.ExternalContact.Name,
				Position:       list.ExternalContact.Position,
				Avatar:         list.ExternalContact.Avatar,
				CorpName:       list.ExternalContact.CorpName,
				CorpFullName:   list.ExternalContact.CorpFullName,
				Type:           list.ExternalContact.Type,
				Gender:         list.ExternalContact.Gender,
				UnionId:        list.ExternalContact.Unionid,
				CreateTime:     utils.GetNow(),
				Owner:          owner,
				TagId:          list.FollowInfo.TagId,
				AddTime:        list.FollowInfo.Createtime,
				Mobiles:        list.FollowInfo.RemarkMobiles,
			}
			if ok {
				follower := &model.CustomerFollowUser{
					Id:                     int64(global.GetSnowflakeNode().Generate()),
					Uid:                    sysUser.Id,
					Userid:                 sysUser.Userid,
					Remark:                 list.FollowInfo.Remark,
					Description:            list.FollowInfo.Description,
					Createtime:             list.FollowInfo.Createtime,
					TagId:                  list.FollowInfo.TagId,
					RemarkCorpName:         list.FollowInfo.RemarkCorpName,
					RemarkMobiles:          list.FollowInfo.RemarkMobiles,
					OperUserid:             list.FollowInfo.OperUserid,
					AddWay:                 list.FollowInfo.AddWay,
					WechatChannelsNickname: list.FollowInfo.WechatChannels.Nickname,
					WechatChannelsSource:   list.FollowInfo.WechatChannels.Source,
					State:                  list.FollowInfo.State,
					FriendState:            model.FriendStateIsFriend,
					ExternalUserId:         list.ExternalContact.ExternalUserid,
				}
				followers = append(followers, follower)
			}

			customers = append(customers, cus)
			counter++
		}
		if len(customers) > 0 {
			if err := customerDao.BatchUpsert(ctx, customers); err != nil {
				log.WithContext(ctx).WithError(err).WithField("customers", customers).Errorf("更新客户信息失败")
				return errors.Wrap(err, "更新客户信息失败")
			}
		}
		if len(followers) > 0 {
			if err := followUserDao.BatchUpsert(ctx, followers); err != nil {
				log.WithContext(ctx).WithError(err).WithField("followers", followers).Errorf("更新好友信息失败")
				return errors.Wrap(err, "更新客户信息失败")
			}
		}
		if cursor == "" {
			break
		}
	}
	log.WithContext(ctx).WithField("userids", userIds).WithField("customers", counter).Infof("同步员工客户完成")
	return nil
}
