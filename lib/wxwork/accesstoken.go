package wxwork

import (
	"fmt"
)

import (
	"context"
)

type AccessTokenResp struct {
	*Response

	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) GetAddressBookToken(ctx context.Context, corpId string, secret string) (string, error) {
	tk, err := c.accessTokenStorage.GetAddressBookToken(ctx, corpId)
	if err != nil {
		resp, err := c.getAddressBookToken(ctx, corpId, secret)
		if err != nil {
			return "", err
		}
		_ = c.accessTokenStorage.SetAddressBookToken(ctx, corpId, resp.AccessToken)
		return resp.AccessToken, nil
	}
	return tk, nil
}

func (c *Client) getAddressBookToken(ctx context.Context, corpId string, secret string) (*AccessTokenResp, error) {
	req := c.newRequest(ctx).SetQueryString(fmt.Sprintf(`corpid=%s&corpsecret=%s`, corpId, secret))
	resp, err := req.Get(accessTokenURL)
	if err != nil {
		return nil, err
	}
	var out AccessTokenResp
	if err := c.jsonDecode(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
