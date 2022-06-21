package dao

import (
	"context"
	"github.com/open-scrm/open-scrm/lib/mongox"
	"github.com/open-scrm/open-scrm/lib/utils"
	"github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDao struct {
	*mongox.MongoX
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{MongoX: mongox.New(model.GetUserColl(ctx))}
}

func (d *UserDao) BasicList(ctx context.Context, deptId []uint32, nameLike string, pageUtil utils.PageUtil) ([]*model.User, int64, error) {
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
	count, err := d.FindAndCount(ctx, query, pageUtil, &out)
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
	opt := options.Find().SetLimit(size).SetSort(bson.M{"_id": 1})
	err := d.Find(ctx, query, &out, opt)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *UserDao) SimpleUserByID(ctx context.Context, ids []int64) (map[int64]*model.SimpleUser, error) {
	query := bson.M{}
	if len(ids) != 0 {
		query["_id"] = mongox.In(ids)
	}
	var out []*model.SimpleUser
	err := d.Find(ctx, query, &out)
	if err != nil {
		return nil, err
	}
	var res = make(map[int64]*model.SimpleUser)
	for _, user := range out {
		res[user.Id] = user
	}
	return res, nil
}
