package mongo

import (
	"context"
	"github.com/open-scrm/open-scrm/internal/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"regexp"
)

type ListParam struct {
	ListRequest *vo.ListReq
	Query       bson.M
}

func List(ctx context.Context, coll *mongo.Collection, param ListParam, res interface{}) (int64, error) {
	cn, err := coll.CountDocuments(ctx, param.Query)
	if err != nil {
		return 0, err
	}
	var opt = options.Find()
	if param.ListRequest.Page > 0 && param.ListRequest.PageSize > 0 {
		opt = opt.SetLimit(param.ListRequest.PageSize).SetSkip((param.ListRequest.Page - 1) * param.ListRequest.PageSize)
	}
	if len(param.ListRequest.Order) != 0 {
		o := bson.M{
			param.ListRequest.Order: 1,
		}
		if !param.ListRequest.Asc {
			o = bson.M{
				param.ListRequest.Order: -1,
			}
		}
		opt = opt.SetSort(o)
	}
	cursor, err := coll.Find(ctx, param.Query)
	if err != nil {
		return 0, err
	}
	err = cursor.All(ctx, res)
	if err != nil {
		return 0, err
	}
	err = cursor.Close(ctx)
	if err != nil {
		return 0, err
	}
	return cn, nil
}

// EscapeText 字符串转义（给正则匹配的特殊字符转译）
func EscapeText(text string) string {
	reg := regexp.MustCompile(`(\.|\?|\*|\+|\\|\$|\^|\(|\)|\|)`)
	return reg.ReplaceAllString(text, "\\${1}")
}

func Like(value string) bson.M {
	return bson.M{"$regex": EscapeText(value), "$options": "i"}
}
