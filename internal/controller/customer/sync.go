package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/service"
	"github.com/open-scrm/open-scrm/pkg/response"
)

func SyncAll(ctx *gin.Context) {
	if err := service.NewSyncExternalContactService().SyncAll(ctx.Request.Context()); err != nil {
		response.SendError(ctx, err)
		return
	}
	response.SendOK(ctx, true)
}
