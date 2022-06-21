package utils

type PageUtil struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"pageSize" form:"pageSize"`
	OrderBy  string `json:"order" form:"order"`
	Asc      bool   `json:"asc" form:"asc"`
}

func (p *PageUtil) SetDefault() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
}
