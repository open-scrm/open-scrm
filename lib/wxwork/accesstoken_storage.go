package wxwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ErrTokenNotFound = errors.New("access token not found")
)

type AccessTokenStorage interface {
	// GetAddressBookToken 获取通讯录管理token
	GetAddressBookToken(ctx context.Context, corpId string) (string, error)
	SetAddressBookToken(ctx context.Context, corpId string, tok string) error

	// GetAppToken 获取应用token
	GetAppToken(ctx context.Context, corpId string, agentId string) (string, error)
	SetAppToken(ctx context.Context, corpId string, agentId string, tok string) error

	// GetExternalContactToken 获取外部联系人 token
	GetExternalContactToken(ctx context.Context, corpId string) (string, error)
	SetExternalContactToken(ctx context.Context, corpId string, tok string) error
}

type redisStorage struct {
	keyPrefix string
	client    *redis.Client
}

func NewRedisStorage(client *redis.Client, keyPrefix string) AccessTokenStorage {
	return &redisStorage{
		keyPrefix: keyPrefix,
		client:    client,
	}
}

func (r *redisStorage) getKey(keyName string) string {
	return fmt.Sprintf("%s.%s", r.keyPrefix, keyName)
}

func (r *redisStorage) GetAddressBookToken(ctx context.Context, corpId string) (string, error) {
	s, _ := r.client.Get(ctx, r.getKey("addressbook."+corpId)).Result()
	if s == "" {
		return "", ErrTokenNotFound
	}
	return s, nil
}

func (r *redisStorage) SetAddressBookToken(ctx context.Context, corpId string, tok string) error {
	return r.client.Set(ctx, r.getKey("addressbook."+corpId), tok, time.Second*7200).Err()
}

func (r *redisStorage) GetAppToken(ctx context.Context, corpId string, agentId string) (string, error) {
	key := r.getKey(fmt.Sprintf("%s.%s.%s", "app", corpId, agentId))
	s, _ := r.client.Get(ctx, key).Result()
	if s == "" {
		return "", ErrTokenNotFound
	}
	return s, nil
}

func (r *redisStorage) SetAppToken(ctx context.Context, corpId string, agentId string, tok string) error {
	key := r.getKey(fmt.Sprintf("%s.%s.%s", "app", corpId, agentId))
	return r.client.Set(ctx, key, tok, time.Second*7200).Err()
}

func (r *redisStorage) GetExternalContactToken(ctx context.Context, corpId string) (string, error) {
	key := r.getKey(fmt.Sprintf("%s.%s", "externalcontact", corpId))
	s, _ := r.client.Get(ctx, key).Result()
	if s == "" {
		return "", ErrTokenNotFound
	}
	return s, nil
}

func (r *redisStorage) SetExternalContactToken(ctx context.Context, corpId string, tok string) error {
	key := r.getKey(fmt.Sprintf("%s.%s", "externalcontact", corpId))
	return r.client.Set(ctx, key, tok, time.Second*7200).Err()
}
