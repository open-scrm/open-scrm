package dao

import (
	"context"
	"github.com/open-scrm/open-scrm/lib/mongox"
	"github.com/open-scrm/open-scrm/pkg/addressbook/model"
	"go.mongodb.org/mongo-driver/bson"
)

type DeptDao struct {
	*mongox.Utilx
}

func NewDeptDao(ctx context.Context) *DeptDao {
	return &DeptDao{Utilx: mongox.New(model.GetDepartmentColl(ctx))}
}

func (d *DeptDao) List(ctx context.Context, ids []uint32, name string) ([]*model.Department, error) {
	var res []*model.Department
	var query = bson.M{}
	if len(ids) != 0 {
		query["_id"] = bson.M{
			"$in": ids,
		}
	}
	if name != "" {
		query["name"] = mongox.Like(name)
	}
	err := d.Find(ctx, query, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *DeptDao) ListMap(ctx context.Context, ids []uint32, name string) (map[uint32]*model.Department, error) {
	res, err := d.List(ctx, ids, name)
	if err != nil {
		return nil, err
	}
	var m = make(map[uint32]*model.Department)
	for _, item := range res {
		m[item.Id] = item
	}
	return m, nil
}
