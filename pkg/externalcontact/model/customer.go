package model

import (
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	Id             int64   `json:"id" bson:"_id"`                        // 主键id
	ExternalUserId string  `json:"externalUserId" bson:"externalUserId"` // 外部联系人id
	Name           string  `json:"name" bson:"name"`
	Position       string  `json:"position" bson:"position"`
	Avatar         string  `json:"avatar" bson:"avatar"`
	CorpName       string  `json:"corpName" bson:"corpName"`
	CorpFullName   string  `json:"corpFullName" bson:"corpFullName"`
	Type           int     `json:"type" bson:"type"`
	Gender         int     `json:"gender" bson:"gender"`
	UnionId        string  `json:"unionId" bson:"unionId"`
	Follows        []int64 `json:"follows" bson:"follows"`

	CreateTime string `json:"createTime" bson:"createTime"`
}

func GetCustomerCollection() *mongo.Collection {
	return global.GetMongoDriver().Database(configs.Get().Mongo.Database).Collection("customer")
}
