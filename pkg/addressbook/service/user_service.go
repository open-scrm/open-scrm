package service

import (
	"context"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/utils"
	"github.com/open-scrm/open-scrm/pkg/addressbook/dao"
	"github.com/open-scrm/open-scrm/pkg/response/addressbook"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) ListUsers(ctx context.Context, req *vo.UserListRequest) (*addressbook.UserListResponse, error) {
	users, count, err := dao.NewUserDao(ctx).BasicList(ctx, req.DeptIds, req.Keyword, req.Page, req.PageSize, req.Order, req.Asc)
	if err != nil {
		log.WithContext(ctx).WithError(err).Errorf("ListUsers错误: %v", req)
		return nil, err
	}
	// 填充部门信息
	var deptIds []uint32
	for _, user := range users {
		deptIds = append(deptIds, user.Department...)
	}
	var res addressbook.UserListResponse
	for _, user := range users {
		res.Data = append(res.Data, &addressbook.UserItem{
			User:  user,
			Depts: nil,
		})
	}
	res.Count = count

	if len(deptIds) != 0 {
		deptIds = utils.SliceSetUint32(deptIds)
		depts, err := dao.NewDeptDao(ctx).ListMap(ctx, deptIds, "")
		if err != nil {
			log.WithContext(ctx).WithError(err).Errorf("查询部门信息失败: %v", req)
			return nil, err
		}
		for _, v := range res.Data {
			for _, dept := range v.Department {
				deptmart, ok := depts[dept]
				if !ok {
					continue
				}
				v.Depts = append(v.Depts, deptmart.Name)
			}
		}
	}

	return &res, err
}
