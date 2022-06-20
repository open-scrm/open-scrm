package model

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id               int64    `json:"id" bson:"_id"`
	Userid           string   `json:"userid" bson:"userid"`
	Name             string   `json:"name" bson:"name"`
	Department       []uint32 `json:"department" bson:"department"`
	Order            []int    `json:"order" bson:"order"`
	Position         string   `json:"position" bson:"position"`
	Mobile           string   `json:"mobile" bson:"mobile"`
	Gender           string   `json:"gender" bson:"gender"`
	Email            string   `json:"email" bson:"email"`
	IsLeaderInDept   []int    `json:"isLeaderInDept" bson:"isLeaderInDept"`
	DirectLeader     []string `json:"directLeader" bson:"directLeader"`
	Avatar           string   `json:"avatar" bson:"avatar"`
	ThumbAvatar      string   `json:"thumbAvatar" bson:"thumbAvatar"`
	Telephone        string   `json:"telephone" bson:"telephone"`
	Alias            string   `json:"alias" bson:"alias"`
	Address          string   `json:"address" bson:"address"`
	OpenUserid       string   `json:"openUserid" bson:"openUserid"`
	MainDepartment   int      `json:"mainDepartment" bson:"mainDepartment"`
	Status           int      `json:"status" bson:"status"`
	QrCode           string   `json:"qrCode" bson:"qrCode"`
	ExternalPosition string   `json:"externalPosition" bson:"externalPosition"`
	CreateTime       string   `json:"createTime" bson:"createTime"`
}

type SimpleUser struct {
	Id         int64    `json:"id" bson:"_id"`
	Userid     string   `json:"userid" bson:"userid"`
	Name       string   `json:"name" bson:"name"`
	Department []uint32 `json:"department" bson:"department"`
}

func GetUserColl(ctx context.Context) *mongo.Collection {
	return global.GetMongoDriver().Database(configs.Get().Mongo.Database).Collection("user")
}
