package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/controller/auth"
	"github.com/open-scrm/open-scrm/lib/session"
)

func authRouter(e *gin.Engine) {
	authGroup := e.Group("/api/v1")
	{
		authGroup.POST("auth/login", auth.Login)
	}

	authdGroup := e.Group("/api/v1", session.Auth())
	{
		authdGroup.GET("/auth/info", auth.UserInfo)
	}
}
