package dao

import (
	"context"
	"github.com/open-scrm/open-scrm/lib/mongox"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerFollowUserDao struct {
	*mongox.MongoX
}

func NewCustomerFollowUserDao() *CustomerFollowUserDao {
	return &CustomerFollowUserDao{
		MongoX: mongox.New(model.GetCustomerFollowUserCollection()),
	}
}

func (d *CustomerFollowUserDao) BatchUpsert(ctx context.Context, list []*model.CustomerFollowUser) error {
	var upsertModel []mongo.WriteModel
	for _, item := range list {
		query := bson.M{
			"uid":            item.Uid,
			"externalUserId": item.ExternalUserId,
		}
		update := bson.M{
			"$set": bson.M{
				"userid":                 item.Userid,
				"remark":                 item.Remark,
				"description":            item.Description,
				"createtime":             item.Createtime,
				"tagId":                  item.TagId,
				"remarkCorpName":         item.RemarkCorpName,
				"remarkMobiles":          item.RemarkMobiles,
				"operUserid":             item.OperUserid,
				"addWay":                 item.AddWay,
				"wechatChannelsNickname": item.WechatChannelsNickname,
				"wechatChannelsSource":   item.WechatChannelsSource,
				"state":                  item.State,
				"friendState":            item.FriendState,
			},
			"$setOnInsert": bson.M{
				"_id":            item.Id,
				"uid":            item.Uid,
				"externalUserId": item.ExternalUserId,
			},
		}
		upsertModel = append(upsertModel, mongo.NewUpdateOneModel().SetFilter(query).SetUpdate(update).SetUpsert(true))
	}
	_, err := d.Coll().BulkWrite(ctx, upsertModel)
	if err != nil {
		return err
	}
	return nil
}
