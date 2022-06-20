package service

import (
	"context"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/utils"
	"github.com/open-scrm/open-scrm/lib/wxwork"
	"github.com/open-scrm/open-scrm/pkg/addressbook/dao"
	dao2 "github.com/open-scrm/open-scrm/pkg/externalcontact/dao"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/model"
	"github.com/open-scrm/open-scrm/pkg/system/service"
	"github.com/pkg/errors"
)

type SyncExternalContactService struct{}

func NewSyncExternalContactService() *SyncExternalContactService {
	return &SyncExternalContactService{}
}

func (s *SyncExternalContactService) SyncAll(ctx context.Context) error {
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
		var ids []string
		for _, user := range users {
			ids = append(ids, user.Userid)
		}
		if len(ids) == 0 {
			break
		}
		if err := s.SyncByUserIds(ctx, ids); err != nil {
			log.WithContext(ctx).WithError(err).WithField("ids", ids).Errorf("同步客户信息失败")
			return err
		}
		lastId = users[len(users)-1].Id
	}
	return nil
}

func (s *SyncExternalContactService) SyncByUserIds(ctx context.Context, userIds []string) error {
	var (
		cursor      string
		limit       int64 = 100
		customerDao       = dao2.NewCustomerDao()
		counter     int64
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

	log.WithContext(ctx).WithField("userids", userIds).Infof("开始同步员工客户")
	for {
		resp, err := cli.BatchGetExternalContactList(ctx, userIds, cursor, limit)
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("批量获取联系人详情失败: %v %v %v", userIds, cursor, limit)
			return err
		}
		cursor = resp.NextCursor
		var customers []*model.Customer
		for _, list := range resp.ExternalContactList {
			// TODO:: 增加好友关联表
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
		if cursor == "" {
			break
		}
	}
	log.WithContext(ctx).WithField("userids", userIds).WithField("customers", counter).Infof("同步员工客户完成")
	return nil
}
