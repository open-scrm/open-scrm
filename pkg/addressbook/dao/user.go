package dao

import (
	"context"
	"github.com/open-scrm/open-scrm/lib/mongox"
	"github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDao struct {
	*mongox.Utilx
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{Utilx: mongox.New(model.GetUserColl(ctx))}
}

func (UserDao) orderBy(orderBy string, asc bool) bson.M {
	sort := -1
	if asc {
		sort = 1
	}
	switch orderBy {
	case "name", "status", "createTime", "_id":
	default:
		orderBy = "_id"
	}
	return bson.M{
		orderBy: sort,
	}
}

func (d *UserDao) BasicList(ctx context.Context, deptId []uint32, nameLike string, page int64, pageSize int64, orderBy string, asc bool) ([]*model.User, int64, error) {
	query := bson.M{}
	if len(deptId) != 0 {
		query["department"] = bson.M{
			"$in": deptId,
		}
	}
	if nameLike != "" {
		query["$or"] = []bson.M{
			{
				"name": mongox.Like(nameLike),
			},
			{
				"mobile": mongox.Like(nameLike),
			},
		}
	}

	var out []*model.User
	opt := options.Find().SetLimit(pageSize).SetSkip((page - 1) * pageSize).SetSort(d.orderBy(orderBy, asc))
	count, err := d.FindAndCount(ctx, query, opt, &out)
	if err != nil {
		return nil, 0, err
	}
	return out, count, nil
}

func (d *UserDao) ScrollList(ctx context.Context, lastId int64, size int64) ([]*model.User, error) {
	query := bson.M{}
	if lastId != 0 {
		query["_id"] = bson.M{
			"$gt": lastId,
		}
	}

	var out []*model.User
	opt := options.Find().SetLimit(size).SetSort(d.orderBy("_id", true))
	err := d.Find(ctx, query, &out, opt)
	if err != nil {
		return nil, err
	}
	return out, nil
}
