package wxwork

import (
	"context"
)

type BatchGetExternalContactListResponse struct {
	Response
	ExternalContactList []ExternalContactList `json:"external_contact_list"`
	NextCursor          string                `json:"next_cursor"`
}

type ExternalContactList struct {
	ExternalContact ExternalContactInfo `json:"external_contact"`
	FollowInfo      ExternalFollowInfo  `json:"follow_info"`
}

type ExternalContactInfo struct {
	ExternalUserid  string `json:"external_userid"`
	Name            string `json:"name"`
	Position        string `json:"position"`
	Avatar          string `json:"avatar"`
	CorpName        string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	Type            int    `json:"type"`
	Gender          int    `json:"gender"`
	Unionid         string `json:"unionid"`
	ExternalProfile struct {
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
	} `json:"external_profile,omitempty"`
}

type ExternalFollowInfo struct {
	Userid         string   `json:"userid"`
	Remark         string   `json:"remark"`
	Description    string   `json:"description"`
	Createtime     int64    `json:"createtime"`
	TagId          []string `json:"tag_id"`
	RemarkCorpName string   `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
	OperUserid     string   `json:"oper_userid"`
	AddWay         int      `json:"add_way"`
	WechatChannels struct {
		Nickname string `json:"nickname"`
		Source   int    `json:"source"`
	} `json:"wechat_channels,omitempty"`
	State string `json:"state,omitempty"`
}

func (c *Client) BatchGetExternalContactList(ctx context.Context, userids []string, cursor string, limit int64) (*BatchGetExternalContactListResponse, error) {
	req := c.newTokenRequest(ctx).SetBody(map[string]interface{}{
		"userid_list": userids,
		"cursor":      cursor,
		"limit":       limit,
	})
	resp, err := req.Post(batchGetExternalContext)
	if err != nil {
		return nil, err
	}
	var out BatchGetExternalContactListResponse
	if err := c.jsonDecode(resp, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
