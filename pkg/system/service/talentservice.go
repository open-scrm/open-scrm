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
)

const (
	redisKeyTalentCache = "cache.talent"
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
			if err := model.GetTalentColl(ctx).FindOne(ctx, bson.M{}).Decode(&res); err != nil {
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
