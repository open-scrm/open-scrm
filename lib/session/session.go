package session

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/pkg/auth/model"
	"github.com/open-scrm/open-scrm/pkg/auth/service"
	"github.com/open-scrm/open-scrm/pkg/response"
)

type contextKey struct{}

var sessionContextKey = contextKey{}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatus(401)
			return
		}
		session, err := service.NewAccountService().GetSession(ctx.Request.Context(), token)
		if err != nil {
			response.SendError(ctx, err)
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), sessionContextKey, session))
		ctx.Next()
	}
}

func GetSession(ctx context.Context) *model.Session {
	return ctx.Value(sessionContextKey).(*model.Session)
}
