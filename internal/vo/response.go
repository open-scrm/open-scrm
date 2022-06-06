package vo

import "github.com/gin-gonic/gin"

type Response struct {
	Data    interface{}
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendOK(ctx *gin.Context, data ...interface{}) {
	var obj interface{}
	if len(data) == 0 {
		obj = struct{}{}
	} else {
		obj = data[0]
	}
	ctx.JSON(200, &Response{
		Data: obj,
	})
}

func SendFail(ctx *gin.Context, msg string) {
	ctx.JSON(400, &Response{
		Code:    400,
		Message: msg,
	})
}

func SendError(ctx *gin.Context, err error) {
	ctx.JSON(500, &Response{
		Code:    500,
		Message: err.Error(), // TODO:: 自定义error
	})
}
