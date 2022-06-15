package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/utils"
	"github.com/open-scrm/open-scrm/pkg/auth/model"
	"time"
)

type AccountService struct{}

func NewAccountService() *AccountService {
	return &AccountService{}
}

// LoginSystem 登录超管.
func (s *AccountService) LoginSystem(ctx context.Context, username string, password string) (string, error) {
	admin, err := defaultAccountProvider.Login(ctx, username, password)
	if err != nil {
		log.WithContext(ctx).Error("登录失败: %v", err)
		return "", err
	}

	// 生成系统token.
	session := model.NewSuperAdminSession(admin)
	// 生成登录日志.
	_, _ = model.GetLoginLogCollection(ctx).InsertOne(ctx, &model.LoginLog{
		Id:          "",
		Username:    admin.Username,
		CreateTime:  utils.GetNow(),
		SessionId:   session.Id,
		UserId:      admin.Id,
		SessionType: session.SessionType,
	})

	data, _ := json.Marshal(session)
	// 生成
	err = global.GetRedis().Set(ctx, session.Id, string(data), time.Duration(configs.Get().Web.SessionExpire)*time.Second).Err()
	if err != nil {
		return "", err
	}
	return session.Id, nil
}

func (s *AccountService) GetSession(ctx context.Context, sid string) (*model.Session, error) {
	var sess model.Session
	data, err := global.GetRedis().Get(ctx, sid).Bytes()
	if err != nil {
		if err != redis.Nil {
			log.WithContext(ctx).WithError(err).Errorf("获取session失败")
			return nil, err
		}
		return nil, err
	}
	err = json.Unmarshal(data, &sess)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

func (s *AccountService) GetSuperAdmin(ctx context.Context, id string) (*model.SuperAdmin, error) {
	return defaultAccountProvider.GetById(ctx, id)
}

func (AccountService) LoginWxWork(ctx context.Context) {

}
