package mapper

import "github.com/open-scrm/open-scrm/pkg/addressbook/model"

type DepartmentTree struct {
	*model.Department
	Children []*DepartmentTree `json:"children"`
}
