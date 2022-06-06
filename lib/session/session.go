package session

import (
	"github.com/gin-gonic/gin"
)

type Session struct {
	Name       string `json:"name"`       // 名称
	UserID     string `json:"userId"`     // 用户id
	IsCorpUser bool   `json:"isCorpUser"` // 是否是企微用户
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
