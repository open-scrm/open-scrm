package model

type SuperAdmin struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}
