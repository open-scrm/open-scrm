package wxwork

import (
	"context"
	"fmt"
)

type ExternalContactListResponse struct {
	Response
	ExternalUserid []string `json:"external_userid"`
}

func (c *Client) ExternalContactList(ctx context.Context, userid string) ([]string, error) {
	req := c.newTokenRequest(ctx).SetQueryString(fmt.Sprintf(`userid=%s`, userid))
	resp, err := req.Get(listExternalContact)
	if err != nil {
		return nil, err
	}
	var out ExternalContactListResponse
	if err := c.jsonDecode(resp, &out); err != nil {
		return nil, err
	}
	return out.ExternalUserid, nil
}
