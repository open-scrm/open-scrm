package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/pkg/externalcontact/service"
	"github.com/open-scrm/open-scrm/pkg/response"
)

func ListCustomer(ctx *gin.Context) {
	req := new(vo.ListCustomerRequest)
	if err := ctx.ShouldBind(req); err != nil {
		response.SendError(ctx, err)
		return
	}
	list, count, err := service.NewCustomerService().List(ctx.Request.Context(), req)
	if err != nil {
		response.SendError(ctx, err)
		return
	}
	response.SendOK(ctx, response.List{
		Data:  list,
		Count: count,
	})
}
