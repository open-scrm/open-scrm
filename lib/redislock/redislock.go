package redislock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	prefix = "lock"
)

type RedisLock struct {
	client *redis.Client
}

func NewRedisLock(client *redis.Client) *RedisLock {
	return &RedisLock{
		client: client,
	}
}

func (l *RedisLock) GetLock(ctx context.Context, key string, expire time.Duration) (bool, error) {
	return l.client.SetNX(ctx, fmt.Sprintf("%s.%s", prefix, key), "1", expire).Result()
}

func (l *RedisLock) UnLock(ctx context.Context, key string) error {
	return l.client.Del(ctx, fmt.Sprintf("%s.%s", prefix, key)).Err()
}
