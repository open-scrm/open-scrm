package model

import "github.com/google/uuid"

const (
	SessionTypeSystem = "sys."
	SessionTypeWxwork = "work."
)

type Session struct {
	Id           string `json:"id"`
	Uid          string `json:"uid"`
	Username     string `json:"username"`
	SessionType  string `json:"sessionType"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	IsSuperAdmin bool   `json:"isSuperAdmin"`
}

func NewSuperAdminSession(sa *SuperAdmin) *Session {
	return &Session{
		Id:           uuid.New().String(),
		Uid:          sa.Id,
		Username:     sa.Username,
		SessionType:  SessionTypeSystem,
		IsSuperAdmin: true,
		Avatar:       sa.Avatar,
	}
}
