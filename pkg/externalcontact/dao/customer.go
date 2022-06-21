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
	*mongox.MongoX
}

func NewCustomerDao() *CustomerDao {
	return &CustomerDao{
		MongoX: mongox.New(model.GetCustomerCollection()),
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
				"remark":       customer.Remark,
				"owner":        customer.Owner,
			},
			"$setOnInsert": bson.M{
				"externalUserId": customer.ExternalUserId,
				"_id":            customer.Id,
				"createTime":     utils.GetNow(),
				"addTime":        customer.AddTime,
			},
			"$addToSet": bson.M{
				"mobiles": bson.M{
					"$each": customer.Mobiles,
				},
				"tagId": bson.M{
					"$each": customer.TagId,
				},
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

type CustomerQuery struct {
	PageUtil utils.PageUtil

	Type          int      // 联系人类型：1微信，2企微
	NameLike      string   // 名称
	RemarkLike    string   // 备注名
	CorpNameLike  string   // 公司名称
	MobileLike    string   // 手机号
	AddWay        int      // 好友来源
	CreateTimeGte int64    // 创建时间大于
	CreateTimeLte int64    // 创建时间小于
	Tags          []string // 企微标签

	// 员工信息
	UserID []int64 // 所属员工
}

func (d *CustomerDao) List(ctx context.Context, query CustomerQuery) ([]*model.Customer, int64, error) {
	q := bson.M{}

	if query.Type != 0 {
		q["type"] = query.Type
	}

	if query.NameLike != "" {
		q["name"] = mongox.Like(query.NameLike)
	}

	if query.RemarkLike != "" {
		q["remark"] = mongox.Like(query.RemarkLike)
	}

	if query.CorpNameLike != "" {
		q["corpFullName"] = mongox.Like(query.CorpNameLike)
	}

	if query.AddWay != 0 {
		q["addWay"] = query.AddWay
	}

	createTimeQuery := bson.M{}
	if query.CreateTimeGte != 0 {
		createTimeQuery["$gte"] = query.CreateTimeGte
	}

	if query.CreateTimeLte != 0 {
		createTimeQuery["$lte"] = query.CreateTimeLte
	}

	if len(createTimeQuery) > 0 {
		createTimeQuery["createtime"] = createTimeQuery
	}

	if len(query.Tags) != 0 {
		q["tagId"] = mongox.In(query.Tags)
	}
	var out []*model.Customer
	count, err := d.FindAndCount(ctx, q, query.PageUtil, &out)
	if err != nil {
		return nil, 0, err
	}
	return out, count, nil
}
