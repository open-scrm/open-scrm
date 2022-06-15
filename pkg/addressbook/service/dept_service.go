package service

import (
	"context"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/mongo"
	"github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"go.mongodb.org/mongo-driver/bson"
)

type DeptService struct{}

func NewDeptService() *DeptService {
	return &DeptService{}
}

// ListDepartments 获取部门列表
// deptIds: 表示获取该级分类下的所有数据.
// 报表数据走MySQL
func (s *DeptService) ListDepartments(ctx context.Context, deptIds []uint32, listReq *vo.ListReq) ([]model.Department, int64, error) {
	q := bson.M{}
	if len(deptIds) != 0 {
		q["_id"] = bson.M{
			"$in": deptIds,
		}
	}

	if listReq.Keyword != "" {
		q["name"] = mongo.Like(listReq.Keyword)
	}

	var res []model.Department
	var count, err = mongo.List(ctx,
		model.GetDepartmentColl(ctx),
		mongo.ListParam{
			ListRequest: listReq,
			Query:       q,
		},
		&res,
	)
	if err != nil {
		log.WithContext(ctx).WithError(err).Errorf("查询部门列表失败")
		return nil, 0, err
	}
	return res, count, nil
}
