package wxwork

import (
	"context"
	"encoding/json"
	"gopkg.in/resty.v1"
)

type Client struct {
	accessTokenStorage AccessTokenStorage

	httpClient *resty.Client
}

func NewClient(accessTokenStorage AccessTokenStorage, httpClient *resty.Client) *Client {
	return &Client{accessTokenStorage: accessTokenStorage, httpClient: httpClient}
}

func (c *Client) r(ctx context.Context) *resty.Request {
	return resty.New().NewRequest().SetContext(ctx)
}

func (c *Client) newRequest(ctx context.Context) *resty.Request {
	return c.r(ctx)
}

func (c *Client) newTokenRequest(ctx context.Context) *resty.Request {
	return c.r(ctx).SetQueryParam("access_token", AccessTokenFromContext(ctx))
}

func (c *Client) jsonDecode(resp *resty.Response, v response) error {
	if err := json.Unmarshal(resp.Body(), v); err != nil {
		return err
	}
	if v.ErrCode() != 0 {
		return &Response{Errmsg: v.ErrMsg(), Errcode: v.ErrCode()}
	}
	return nil
}

var DefaultHttpClient = resty.New()
