package service

import (
	"context"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/mongox"
	"github.com/open-scrm/open-scrm/pkg/addressbook/dao"
	"github.com/open-scrm/open-scrm/pkg/addressbook/mapper"
	"github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"go.mongodb.org/mongo-driver/bson"
)

type DeptService struct{}

func NewDeptService() *DeptService {
	return &DeptService{}
}

// ListDepartmentTree 获取部门列表
// deptIds: 表示获取该级分类下的所有数据. 返回树形结构.
func (s *DeptService) ListDepartmentTree(ctx context.Context, deptIds []uint32, name string) (*mapper.DepartmentTree, error) {
	d := dao.NewDeptDao(ctx)
	res, err := d.List(ctx, deptIds, name)
	if err != nil {
		return nil, err
	}
	return toTree(res, model.RootDeptId), nil
}

func toTree(data []*model.Department, pid uint32) *mapper.DepartmentTree {
	t := &mapper.DepartmentTree{
		Department: nil,
		Children:   nil,
	}
	for _, d := range data {
		if d.Id == pid {
			t.Department = d
		} else {
			if d.ParentId == pid {
				t.Children = append(t.Children, toTree(data, d.Id))
			}
		}
	}
	return t
}

// ListDepartments 获取部门列表
// deptIds: 表示获取该级分类下的所有数据.
// 报表数据走MySQL
func (s *DeptService) ListDepartments(ctx context.Context, deptIds []uint32, name string) ([]*model.Department, int64, error) {
	q := bson.M{}
	if len(deptIds) != 0 {
		q["_id"] = bson.M{
			"$in": deptIds,
		}
	}

	if name != "" {
		q["name"] = mongox.Like(name)
	}

	var res []*model.Department
	var count, err = mongox.List(ctx,
		model.GetDepartmentColl(ctx),
		mongox.ListParam{
			Query: q,
		},
		&res,
	)
	if err != nil {
		log.WithContext(ctx).WithError(err).Errorf("查询部门列表失败")
		return nil, 0, err
	}
	return res, count, nil
}
