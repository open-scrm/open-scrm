package vo

type ListReq struct {
	Page     int64  `json:"page" binding:"required"`
	PageSize int64  `json:"pageSize" binding:"required"`
	Order    string `json:"order"`
	Asc      bool   `json:"asc"`
	Keyword  string `json:"keyword"`
}

func (r *ListReq) SetDefault() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 20
	}
}

// AddressBookListDeptRequest 查询所有的分组并返回树形结构
type AddressBookListDeptRequest struct {
	Name string `json:"name"`
}

func (r *AddressBookListDeptRequest) Validate() error {
	return nil
}

// UserListRequest 员工列表
type UserListRequest struct {
	ListReq
	DeptIds []uint32 `json:"deptIds"`
}

func (r *UserListRequest) Validate() error {
	r.SetDefault()
	return nil
}
