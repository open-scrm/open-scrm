package addressbook

import "github.com/open-scrm/open-scrm/pkg/addressbook/model"

type UserListResponse struct {
	Data  []*UserItem `json:"data"`
	Count int64       `json:"count"`
}

type UserItem struct {
	*model.User
	Roles []string `json:"roles" description:"角色"`
	Depts []string `json:"depts" description:"部门"`
}
