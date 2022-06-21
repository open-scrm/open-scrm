package service

import (
	"context"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/utils"
	dao2 "github.com/open-scrm/open-scrm/pkg/addressbook/dao"
	model2 "github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/dao"
	"github.com/open-scrm/open-scrm/pkg/response/externalcontact"
)

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) List(ctx context.Context, req *vo.ListCustomerRequest) ([]*externalcontact.CustomerListResponse, int64, error) {
	d := dao.NewCustomerDao()
	list, count, err := d.List(ctx, dao.CustomerQuery{
		PageUtil:      req.PageUtil,
		Type:          req.Type,
		NameLike:      req.Name,
		RemarkLike:    req.Remark,
		CorpNameLike:  req.CorpName,
		MobileLike:    req.Mobile,
		AddWay:        req.AddWay,
		CreateTimeGte: req.CreateTimeGte,
		CreateTimeLte: req.CreateTimeLte,
		Tags:          req.Tags,
		//UserID:        nil,
	})
	if err != nil {
		log.ErrCtx(ctx, err).Errorf("查询客户信息失败: %v", req)
		return nil, 0, err
	}

	var (
		userMap = make(map[int64]*model2.SimpleUser)
	)
	// 查询员工信息
	var userIds []int64
	for _, user := range list {
		userIds = append(userIds, user.Owner)
	}
	if len(userIds) > 0 {
		userDao := dao2.NewUserDao(ctx)
		users, err := userDao.SimpleUserByID(ctx, userIds)
		if err != nil {
			log.ErrCtx(ctx, err).Errorf("查询员工信息失败: %v", userIds)
			return nil, 0, err
		}
		userMap = users
	}

	var res []*externalcontact.CustomerListResponse
	for _, item := range list {
		res = append(res, &externalcontact.CustomerListResponse{
			Customer:  item,
			OwnerInfo: userMap[item.Owner],
			AddTime:   utils.FromUnix(item.AddTime),
		})
	}
	return res, count, nil
}
