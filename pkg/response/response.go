package response

import "github.com/gin-gonic/gin"

type Response struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Stack   interface{} `json:"stack"`
}

type List struct {
	Data  interface{} `json:"data"`
	Count interface{} `json:"count"`
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
	var res = Response{
		Code:    500,
		Message: err.Error(), // TODO:: 自定义error
	}
	if e, ok := err.(*Error); ok {
		res.Code = e.code
		res.Message = e.msg
		if gin.IsDebugging() {
			res.Stack = e
		}
	}
	ctx.JSON(500, res)
}
