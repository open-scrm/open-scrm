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
	if r.PageSize <= 20 {
		r.PageSize = 20
	}
}

// AddressBookListDeptRequest 查询所有的分组并返回树形结构
type AddressBookListDeptRequest struct {
	*ListReq
}

func (r *AddressBookListDeptRequest) Validate() error {
	r.SetDefault()
	r.Page = 0
	r.PageSize = 0
	return nil
}
