package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/pkg/system/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	redisKeyTalentCache = "cache.talent"
	defaultTalentId     = "defaultTalentId"
)

type TalentService struct{}

func NewTalentService() *TalentService {
	return &TalentService{}
}

func (t *TalentService) GetTalentInfo(ctx context.Context) (model.Talent, error) {
	data, err := global.GetRedis().Get(ctx, redisKeyTalentCache).Bytes()
	if err != nil {
		if err == redis.Nil {
			var res model.Talent
			if err := model.GetTalentColl(ctx).FindOne(ctx, bson.M{"_id": defaultTalentId}).Decode(&res); err != nil {
				if err == mongo.ErrNoDocuments {
					return model.Talent{}, vo.NewError(ctx, "租户配置信息不存在")
				}
				return model.Talent{}, err
			}
			data, _ := json.Marshal(res)
			if err := global.GetRedis().Set(ctx, redisKeyTalentCache, string(data), 0).Err(); err != nil {
				log.WithContext(ctx).WithError(err).WithField("key", redisKeyTalentCache).Errorf("写入redis缓存失败")
			}

			return res, nil
		}
		return model.Talent{}, err
	}

	var talent model.Talent
	if err := json.Unmarshal(data, &talent); err != nil {
		return model.Talent{}, nil
	}
	return talent, nil
}

/**

type Talent struct {
	Id                    string `json:"id" bson:"_id"`
	CorpId                string `json:"corpId" bson:"corpId"`
	AgentId               string `json:"agentId" bson:"agentId"`
	Db                    string `json:"db" bson:"db"`
	AddressBookSecret     string `json:"addressBookSecret" bson:"addressBookSecret"`
	AppSecret             string `json:"appSecret" bson:"appSecret"`
	ExternalContactSecret string `json:"externalContactSecret" bson:"externalContactSecret"`
}
*/

// SetTalentInfo 设置 talent 信息.
// id 为固定的值.
func (t *TalentService) SetTalentInfo(ctx context.Context, talent model.Talent) error {
	query := bson.M{
		"_id": defaultTalentId,
	}
	update := bson.M{
		"$set": bson.M{
			"corpId":                talent.CorpId,
			"agentId":               talent.AgentId,
			"db":                    talent.Db,
			"addressBookSecret":     talent.AddressBookSecret,
			"appSecret":             talent.AppSecret,
			"externalContactSecret": talent.AppSecret,
		},
		"$setOnInsert": bson.M{
			"_id": defaultTalentId,
		},
	}
	_, err := model.GetTalentColl(ctx).UpdateOne(ctx, query, update, options.Update().SetUpsert(true))

	// 删除redis 缓存.
	global.GetRedis().Del(ctx, redisKeyTalentCache)
	return err
}
