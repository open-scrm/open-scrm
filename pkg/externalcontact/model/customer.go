package model

import (
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	Id             int64    `json:"id" bson:"_id"`                        // 主键id
	ExternalUserId string   `json:"externalUserId" bson:"externalUserId"` // 外部联系人id
	Name           string   `json:"name" bson:"name"`                     // 名称
	Position       string   `json:"position" bson:"position"`             // 职位
	Avatar         string   `json:"avatar" bson:"avatar"`                 // 头像
	CorpName       string   `json:"corpName" bson:"corpName"`             // 企业
	CorpFullName   string   `json:"corpFullName" bson:"corpFullName"`     // 企业
	Type           int      `json:"type" bson:"type"`                     // 1企微. 2微信
	Gender         int      `json:"gender" bson:"gender"`                 // 性别
	UnionId        string   `json:"unionId" bson:"unionId"`               // unionId
	Follows        []int64  `json:"follows" bson:"follows"`               // 共享人
	CreateTime     string   `json:"createTime" bson:"createTime"`         // 创建时间
	Owner          int64    `json:"owner" bson:"owner"`                   // 跟进人
	Remark         string   `json:"remark" bson:"remark"`                 // 备注名
	TagId          []string `json:"tagId" bson:"tagId"`                   // 该成员添加此外部联系人所打企业标签的id，用户自定义类型标签（type=2）不返回
	AddTime        int64    `json:"addTime" bson:"addTime"`               // 添加时间
	Mobiles        []string `json:"mobiles" bson:"mobiles"`               // 电话

	// 基本信息
	WechatNo string `json:"wechatNo" bson:"wechatNo"` // 微信号
	Age      int    `json:"age" bson:"age"`           // 年龄
	Birthday string `json:"birthday" bson:"birthday"` // 生日
	Address  string `json:"address" bson:"address"`   // 地址
	Province string `json:"province" bson:"province"` // 省
	City     string `json:"city" bson:"city"`         // 市
	Area     string `json:"area" bson:"area"`         // 区
}

func GetCustomerCollection() *mongo.Collection {
	return global.GetMongoDriver().Database(configs.Get().Mongo.Database).Collection("customer")
}
