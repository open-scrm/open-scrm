package wxwork

import (
	"context"
	"fmt"
)

type DepartmentListResponse struct {
	*Response

	Department []struct {
		Id               uint32   `json:"id"`
		Name             string   `json:"name"`
		NameEn           string   `json:"name_en"`
		DepartmentLeader []string `json:"department_leader"`
		Parentid         int      `json:"parentid"`
		Order            int      `json:"order"`
	} `json:"department"`
}

func (c *Client) ListDepartment(ctx context.Context, id string) (*DepartmentListResponse, error) {
	req := c.newTokenRequest(ctx).SetQueryString(fmt.Sprintf(`id=%s`, id))
	resp, err := req.Get(departmentList)
	if err != nil {
		return nil, err
	}
	var out DepartmentListResponse
	if err := c.jsonDecode(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
