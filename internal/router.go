package internal

import (
	"github.com/clearcodecn/swaggos"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/controller/auth"
	"github.com/open-scrm/open-scrm/internal/controller/auth/vo"
	"github.com/open-scrm/open-scrm/lib/session"
	"github.com/open-scrm/open-scrm/pkg/response"
)

func authRouter(e *gin.Engine, doc *swaggos.Swaggos) {
	authGroup := e.Group("/api/v1")
	docAuthGroup := doc.Group("/api/v1").Tag("auth")
	{
		authGroup.POST("auth/login", auth.Login)
		docAuthGroup.Post("auth/login").Body(new(vo.LoginRequest)).JSON(response.NewResponse("session id")).Description("登录接口")
	}

	authdGroup := e.Group("/api/v1", session.Auth())
	{
		authdGroup.GET("/auth/info", auth.UserInfo)
		docAuthGroup.Get("auth/info").JSON(response.NewResponse(&vo.AuthInfoResponse{})).Description("获取当前账号登录信息")
	}
}
