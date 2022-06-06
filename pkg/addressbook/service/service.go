package service

import (
	"context"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/open-scrm/open-scrm/lib/wxwork"
	"github.com/open-scrm/open-scrm/pkg/system/service"
	"github.com/pkg/errors"
)

type SyncCorpStructureService struct{}

func NewSyncCorpStructureService() *SyncCorpStructureService {
	return &SyncCorpStructureService{}
}

func (s *SyncCorpStructureService) DoSync(ctx context.Context) error {
	talentInfo, err := service.NewTalentService().GetTalentInfo(ctx)
	if err != nil {
		return errors.Wrap(err, "获取租户信息失败")
	}
	cli := global.GetWxWorkClient()

	token, err := cli.GetAddressBookToken(ctx, talentInfo.CorpId, talentInfo.AddressBookSecret)
	if err != nil {
		return errors.Wrap(err, "获取access token失败")
	}
	ctx = wxwork.ContextWithAccessToken(ctx, token)
	department, err := global.GetWxWorkClient().ListDepartment(ctx, "")
	if err != nil {
		return errors.Wrap(err, "获取部门列表失败")
	}
	for _, item := range department.Department {
		_ = item
	}
	return nil
}
