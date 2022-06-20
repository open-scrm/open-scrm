package vo

import (
	"github.com/open-scrm/open-scrm/pkg/response"
)

type LoginRequest struct {
	Username string `json:"username" description:"账号，开发环境是 admin" required:"true"`
	Password string `json:"password" description:"密码: 开发环境是 admin" required:"true"`
}

func (r *LoginRequest) Validate() error {
	if r.Username == "" || r.Password == "" {
		return response.InvalidParam(nil)
	}
	return nil
}

type AuthInfoResponse struct {
	Id          string   `json:"id"`
	Username    string   `json:"username"`
	Avatar      string   `json:"avatar"`
	Nickname    string   `json:"nickname"`
	Permissions []string `json:"permissions"`
}
