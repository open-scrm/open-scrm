package model

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type Talent struct {
	Id                    string `json:"id" bson:"_id"`
	CorpId                string `json:"corpId" bson:"corpId"`
	AgentId               string `json:"agentId" bson:"agentId"`
	Db                    string `json:"db" bson:"db"`
	AddressBookSecret     string `json:"addressBookSecret" bson:"addressBookSecret"`
	AppSecret             string `json:"appSecret" bson:"appSecret"`
	ExternalContactSecret string `json:"externalContactSecret" bson:"externalContactSecret"`
}

func GetTalentColl(ctx context.Context) *mongo.Collection {
	return global.GetMongoDriver().Database(configs.Get().Mongo.Database).Collection("talent")
}
