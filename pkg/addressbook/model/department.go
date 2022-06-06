package model

import "context"

type Department struct {
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
	ParentId int    `json:"parentid"`
	Order    int    `json:"order"`
	Id       int    `json:"id"`
}

func DepartmentCollection(ctx context.Context) string {
	return "department"
}
