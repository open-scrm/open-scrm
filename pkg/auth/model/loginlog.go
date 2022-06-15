package model

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginLog struct {
	Id          string `json:"id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	CreateTime  string `json:"createTime" bson:"createTime"`
	SessionId   string `json:"sessionId" bson:"sessionId"`
	UserId      string `json:"userId" bson:"userId"`
	SessionType string `json:"sessionType" bson:"sessionType"`
	Ip          string `json:"ip" bson:"ip"`
}

func GetLoginLogCollection(ctx context.Context) *mongo.Collection {
	return global.GetMongoDriver().Database(configs.Get().Mongo.Database).Collection("login_log")
}
