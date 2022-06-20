package mongox

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
)

type ListParam struct {
	Query bson.M
}

func List(ctx context.Context, coll *mongo.Collection, param ListParam, res interface{}) (int64, error) {
	cn, err := coll.CountDocuments(ctx, param.Query)
	if err != nil {
		return 0, err
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
