package vo

import "github.com/open-scrm/open-scrm/lib/utils"

// AddressBookListDeptRequest 查询所有的分组并返回树形结构
type AddressBookListDeptRequest struct {
	Name string `json:"name"`
}

func (r *AddressBookListDeptRequest) Validate() error {
	return nil
}

// UserListRequest 员工列表
type UserListRequest struct {
	utils.PageUtil
	Keyword string   `json:"keyword"`
	DeptIds []uint32 `json:"deptIds"`
}

func (r *UserListRequest) Validate() error {
	r.SetDefault()
	return nil
}
