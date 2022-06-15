package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/controller/auth/vo"
	"github.com/open-scrm/open-scrm/lib/session"
	"github.com/open-scrm/open-scrm/pkg/auth/model"
	"github.com/open-scrm/open-scrm/pkg/auth/service"
	"github.com/open-scrm/open-scrm/pkg/response"
)

func Login(ctx *gin.Context) {
	req := new(vo.LoginRequest)
	if err := ctx.ShouldBind(req); err != nil {
		response.SendError(ctx, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.SendError(ctx, err)
		return
	}

	s := service.NewAccountService()
	sessionId, err := s.LoginSystem(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		response.SendError(ctx, err)
		return
	}
	response.SendOK(ctx, sessionId)
}

func UserInfo(ctx *gin.Context) {
	sess := session.GetSession(ctx.Request.Context())
	switch sess.SessionType {
	case model.SessionTypeSystem:
		info, err := service.NewAccountService().GetSuperAdmin(ctx.Request.Context(), sess.Uid)
		if err != nil {
			response.SendError(ctx, err)
			return
		}
		response.SendOK(ctx, &vo.AuthInfoResponse{
			Id:          info.Id,
			Username:    info.Username,
			Avatar:      info.Avatar,
			Nickname:    info.Nickname,
			Permissions: []string{"*"},
		})
		return
	}
}
