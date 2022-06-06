package log

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func Start(ctx context.Context, name string) opentracing.Span {
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		return opentracing.StartSpan(name)
	} else {
		return opentracing.StartSpan(name, opentracing.ChildOf(sp.Context()))
	}
}
