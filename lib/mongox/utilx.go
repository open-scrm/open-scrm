package mongox

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Utilx struct {
	dao *mongo.Collection
}

func New(dao *mongo.Collection) *Utilx {
	return &Utilx{dao: dao}
}

func (d *Utilx) Coll() *mongo.Collection {
	return d.dao
}

func (d *Utilx) FindAndCount(ctx context.Context, query bson.M, options *options.FindOptions, out interface{}) (int64, error) {
	cn, err := d.Coll().CountDocuments(ctx, query)
	if err != nil {
		return 0, err
	}

	cursor, err := d.Coll().Find(ctx, query, options)
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

func (d *Utilx) Find(ctx context.Context, query bson.M, out interface{}, options ...*options.FindOptions) error {
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
