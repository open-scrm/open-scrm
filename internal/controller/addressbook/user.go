package addressbook

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/pkg/addressbook/service"
	"github.com/open-scrm/open-scrm/pkg/response"
)

func UserList(ctx *gin.Context) {
	req := new(vo.UserListRequest)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithContext(ctx.Request.Context()).WithError(err).Errorf("UserListRequest 参数错误")
		response.SendFail(ctx, "参数错误")
		return
	}
	if err := req.Validate(); err != nil {
		response.SendError(ctx, err)
		return
	}
	list, err := service.NewUserService().ListUsers(ctx, req)
	if err != nil {
		log.WithContext(ctx.Request.Context()).WithError(err).Errorf("DepartmentList 查询列表失败")
		response.SendError(ctx, err)
		return
	}
	response.SendOK(ctx, list)
}
