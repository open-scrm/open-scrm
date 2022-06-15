package vo

import (
	"github.com/open-scrm/open-scrm/pkg/response"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
