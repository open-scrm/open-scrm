package service

import (
	"context"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/utils"
	"github.com/open-scrm/open-scrm/lib/wxwork"
	"github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"github.com/open-scrm/open-scrm/pkg/system/service"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SyncCorpStructureService struct{}

func NewSyncCorpStructureService() *SyncCorpStructureService {
	return &SyncCorpStructureService{}
}

func (s *SyncCorpStructureService) DoSync(ctx context.Context) error {
	sp := log.Start(ctx, "syncCorp")
	defer sp.Finish()

	log.WithContext(ctx).Infof("开始同步组织架构")

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

	if err := s.syncDepartment(ctx); err != nil {
		return err
	}
	if err := s.syncUsers(ctx); err != nil {
		return err
	}
	log.WithContext(ctx).Infof("同步组织架构成功")
	return nil
}

func (s *SyncCorpStructureService) syncDepartment(ctx context.Context) error {
	department, err := global.GetWxWorkClient().ListDepartment(ctx, "")
	if err != nil {
		log.WithContext(ctx).WithError(err).Errorf("获取部门列表失败")
		return errors.Wrap(err, "获取部门列表失败")
	}
	var (
		depts   []model.Department
		deptIds []uint32
	)
	for _, item := range department.Department {
		dept := model.Department{
			Name:     item.Name,
			NameEn:   item.NameEn,
			ParentId: item.Parentid,
			Order:    item.Order,
			Id:       item.Id,
		}
		depts = append(depts, dept)
		deptIds = append(deptIds, item.Id)
	}

	dao := model.GetDepartmentColl(ctx)
	if len(deptIds) > 0 {
		_, err = dao.DeleteMany(ctx, bson.M{
			"_id": bson.M{
				"$nin": deptIds,
			},
		})
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("删除部门失败")
			return errors.Wrap(err, "删除部门失败")
		}
	}

	var writeModels []mongo.WriteModel
	for _, item := range depts {
		update := bson.M{
			"$set": bson.M{
				"name":     item.Name,
				"nameEn":   item.NameEn,
				"parentId": item.ParentId,
				"order":    item.Order,
			},
			"$setOnInsert": bson.M{
				"_id": item.Id,
			},
		}
		writeModels = append(writeModels,
			mongo.
				NewUpdateOneModel().
				SetUpsert(true).
				SetFilter(bson.M{"_id": item.Id}).
				SetUpdate(update),
		)
	}
	// 批量更新.
	if _, err := dao.BulkWrite(ctx, writeModels); err != nil {
		log.WithContext(ctx).WithError(err).Errorf("批量更新部门失败")
		return errors.Wrap(err, "批量更新部门失败")
	}
	return nil
}

func (s *SyncCorpStructureService) syncUsers(ctx context.Context) error {
	users, err := global.GetWxWorkClient().ListUserByDept(ctx, 1, true)
	if err != nil {
		log.WithContext(ctx).WithError(err).Errorf("获取员工列表失败")
		return err
	}
	var writeModels []mongo.WriteModel
	for _, user := range users.Userlist {
		resp, err := global.GetWxWorkClient().GetUserDetail(ctx, user.Userid)
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("获取员工详情失败")
			return err
		}
		m := mongo.NewUpdateOneModel().SetFilter(bson.M{"userid": resp.Userid}).SetUpsert(true)
		m.SetUpdate(bson.M{
			"$set": bson.M{
				"name":             resp.Name,
				"department":       resp.Department,
				"order":            resp.Order,
				"position":         resp.Position,
				"mobile":           resp.Mobile,
				"gender":           resp.Gender,
				"email":            resp.Email,
				"isLeaderInDept":   resp.IsLeaderInDept,
				"directLeader":     resp.DirectLeader,
				"avatar":           resp.Avatar,
				"thumbAvatar":      resp.ThumbAvatar,
				"telephone":        resp.Telephone,
				"alias":            resp.Alias,
				"address":          resp.Address,
				"openUserid":       resp.OpenUserid,
				"mainDepartment":   resp.MainDepartment,
				"status":           resp.Status,
				"qrCode":           resp.QrCode,
				"externalPosition": resp.ExternalPosition,
			},
			"$setOnInsert": bson.M{
				"_id":        global.GetSnowflakeNode().Generate(),
				"userid":     resp.Userid,
				"createTime": utils.GetNow(),
			},
		})
		writeModels = append(writeModels, m)
	}

	if len(writeModels) > 0 {
		dao := model.GetUserColl(ctx)
		if _, err := dao.BulkWrite(ctx, writeModels); err != nil {
			log.WithContext(ctx).WithError(err).Errorf("批量更新员工失败")
			return err
		}
	}
	return nil
}
