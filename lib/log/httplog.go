package log

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Logger() gin.HandlerFunc {
	return gin.Logger()
	//return func(ctx *gin.Context) {
	//	// TODO:: logger.
	//	ctx.Next()
	//}
}

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sp := opentracing.StartSpan("request")
		defer sp.Finish()

		sp.SetTag("http.method", ctx.Request.Method)
		sp.SetTag("http.url", ctx.Request.URL.String())

		ctx.Request = ctx.Request.WithContext(opentracing.ContextWithSpan(ctx.Request.Context(), sp))
		ctx.Header("x-request-id", sp.Context().(jaeger.SpanContext).TraceID().String())
		ctx.Next()

		sp.SetTag("http.statusCode", ctx.Writer.Status())
		sp.SetTag("http.size", ctx.Writer.Size())
	}
}
