package dao

import (
	"context"
	"github.com/open-scrm/open-scrm/lib/mongox"
	"github.com/open-scrm/open-scrm/lib/utils"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerDao struct {
	*mongox.Utilx
}

func NewCustomerDao() *CustomerDao {
	return &CustomerDao{
		Utilx: mongox.New(model.GetCustomerCollection()),
	}
}

func (d *CustomerDao) BatchUpsert(ctx context.Context, customers []*model.Customer) error {
	var upsertModel []mongo.WriteModel
	for _, customer := range customers {
		query := bson.M{
			"externalUserId": customer.ExternalUserId,
		}
		update := bson.M{
			"$set": bson.M{
				"name":         customer.Name,
				"position":     customer.Position,
				"avatar":       customer.Avatar,
				"corpName":     customer.CorpName,
				"corpFullName": customer.CorpFullName,
				"type":         customer.Type,
				"gender":       customer.Gender,
				"unionId":      customer.UnionId,
			},
			"$setOnInsert": bson.M{
				"externalUserId": customer.ExternalUserId,
				"_id":            customer.Id,
				"createTime":     utils.GetNow(),
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
