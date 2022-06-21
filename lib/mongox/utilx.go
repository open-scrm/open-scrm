package mongox

import (
	"context"
	"github.com/open-scrm/open-scrm/lib/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoX struct {
	dao *mongo.Collection
}

func New(dao *mongo.Collection) *MongoX {
	return &MongoX{dao: dao}
}

func (d *MongoX) Coll() *mongo.Collection {
	return d.dao
}

func (MongoX) orderBy(orderBy string, asc bool) bson.M {
	sort := -1
	if asc {
		sort = 1
	}
	if orderBy == "" {
		orderBy = "_id"
	}
	return bson.M{
		orderBy: sort,
	}
}

func (d *MongoX) FindAndCount(ctx context.Context, query bson.M, pageUtil utils.PageUtil, out interface{}) (int64, error) {
	cn, err := d.Coll().CountDocuments(ctx, query)
	if err != nil {
		return 0, err
	}

	opt := options.Find().SetLimit(pageUtil.PageSize).SetSkip((pageUtil.Page - 1) * pageUtil.PageSize).SetSort(d.orderBy(pageUtil.OrderBy, pageUtil.Asc))
	cursor, err := d.Coll().Find(ctx, query, opt)
	if err != nil {
		return 0, err
	}
	err = cursor.All(ctx, out)
	if err != nil {
		return 0, err
	}
	err = cursor.Close(ctx)
	if err != nil {
		return 0, err
	}
	return cn, nil
}

func (d *MongoX) Find(ctx context.Context, query bson.M, out interface{}, options ...*options.FindOptions) error {
	cursor, err := d.Coll().Find(ctx, query, options...)
	if err != nil {
		return err
	}
	err = cursor.All(ctx, out)
	if err != nil {
		_ = cursor.Close(ctx)
		return err
	}
	err = cursor.Close(ctx)
	if err != nil {
		return err
	}
	return nil
}
