package service

import (
	"context"
	"fmt"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/pkg/auth/model"
	"github.com/open-scrm/open-scrm/pkg/response"
)

var (
	defaultAccountProvider = newConfigFileAccountProvider()
)

type AccountProvider interface {
	GetById(ctx context.Context, id string) (*model.SuperAdmin, error)
	Login(ctx context.Context, username string, password string) (*model.SuperAdmin, error)
}

type configFileAccountProvider struct {
}

func newConfigFileAccountProvider() *configFileAccountProvider {
	return &configFileAccountProvider{}
}

func (c *configFileAccountProvider) GetById(ctx context.Context, id string) (*model.SuperAdmin, error) {
	sa := configs.Get().SuperAdmin
	for _, admin := range sa {
		if id == admin.Id {
			return c.convertConfigToModel(admin), nil
		}
	}
	return nil, fmt.Errorf("user not exist")
}

func (c configFileAccountProvider) convertConfigToModel(admin configs.AdminConfig) *model.SuperAdmin {
	return &model.SuperAdmin{
		Id:       admin.Id,
		Username: admin.Username,
		Password: admin.Password,
		Avatar:   admin.Avatar,
		Nickname: admin.Nickname,
	}
}

func (c *configFileAccountProvider) Login(ctx context.Context, username string, password string) (*model.SuperAdmin, error) {
	sa := configs.Get().SuperAdmin
	for _, admin := range sa {
		if admin.Username == username && admin.Password == password {
			return c.convertConfigToModel(admin), nil
		}
	}
	return nil, response.NewError(ctx, "用户名或密码错误")
}
