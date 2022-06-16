package model

import (
	"context"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	RootDeptId = 1
)

type Department struct {
	Name     string `json:"name" bson:"name"`
	NameEn   string `json:"nameEn" bson:"nameEn"`
	ParentId uint32 `json:"parentId" bson:"parentId"`
	Order    uint32 `json:"order" bson:"order"`
	Id       uint32 `json:"id" bson:"_id"`
}

func GetDepartmentColl(ctx context.Context) *mongo.Collection {
	return global.GetMongoDriver().Database(configs.Get().Mongo.Database).Collection("department")
}
