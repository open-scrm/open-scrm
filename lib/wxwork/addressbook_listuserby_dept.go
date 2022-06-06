package wxwork

import (
	"context"
	"fmt"
)

type ListUserResponse struct {
	*Response
	Userlist []struct {
		Userid     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserid string `json:"open_userid"`
	} `json:"userlist"`
}

func (c *Client) ListUserByDept(ctx context.Context, deptId uint32, fetchChildren bool) (*ListUserResponse, error) {
	var loop = "0"
	if fetchChildren {
		loop = "1"
	}
	req := c.newTokenRequest(ctx).SetQueryString(fmt.Sprintf(`department_id=%d&fetch_child=%s`, deptId, loop))
	resp, err := req.Get(userListByDept)
	if err != nil {
		return nil, err
	}
	var out ListUserResponse
	if err := c.jsonDecode(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type GetUserDetailResponse struct {
	*Response

	Userid         string   `json:"userid"`
	Name           string   `json:"name"`
	Department     []int    `json:"department"`
	Order          []int    `json:"order"`
	Position       string   `json:"position"`
	Mobile         string   `json:"mobile"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	BizMail        string   `json:"biz_mail"`
	IsLeaderInDept []int    `json:"is_leader_in_dept"`
	DirectLeader   []string `json:"direct_leader"`
	Avatar         string   `json:"avatar"`
	ThumbAvatar    string   `json:"thumb_avatar"`
	Telephone      string   `json:"telephone"`
	Alias          string   `json:"alias"`
	Address        string   `json:"address"`
	OpenUserid     string   `json:"open_userid"`
	MainDepartment int      `json:"main_department"`
	Extattr        struct {
		Attrs []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
		} `json:"attrs"`
	} `json:"extattr"`
	Status           int    `json:"status"`
	QrCode           string `json:"qr_code"`
	ExternalPosition string `json:"external_position"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name"`
		WechatChannels   struct {
			Nickname string `json:"nickname"`
			Status   int    `json:"status"`
		} `json:"wechat_channels"`
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		} `json:"external_attr"`
	} `json:"external_profile"`
}

func (c *Client) GetUserDetail(ctx context.Context, userId string) (*GetUserDetailResponse, error) {
	req := c.newTokenRequest(ctx).SetQueryString(fmt.Sprintf(`userid=%s`, userId))
	resp, err := req.Get(getUser)
	if err != nil {
		return nil, err
	}
	var out GetUserDetailResponse
	if err := c.jsonDecode(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
