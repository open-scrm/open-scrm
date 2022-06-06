package log

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"io"
	"path/filepath"
	"runtime"
	"strings"
)

type callerFormat struct {
	children logrus.Formatter
}

func (c *callerFormat) getCurrentPosition(entry *logrus.Entry) (string, string, int) {
	skip := 4
	if len(entry.Data) == 0 {
		skip = 6
	}
start:
	pc, file, line, _ := runtime.Caller(skip)
	function := runtime.FuncForPC(pc).Name()
	if strings.LastIndex(function, "sirupsen/logrus.") >= 0 ||
		strings.LastIndex(function, "log/logrusx.") >= 0 {
		skip++
		goto start
	}
	function = strings.TrimPrefix(function, "github.com/open-scrm/open-scrm/")
	return function, file, line
}

func (c *callerFormat) Format(e *logrus.Entry) ([]byte, error) {
	function, file, line := c.getCurrentPosition(e)
	caller := fmt.Sprintf("%s[%s:%d]", function, filepath.Base(file), line)
	if _, ok := e.Data["file"]; !ok {
		e.Data["file"] = caller
	}
	return c.children.Format(e)
}

type tracingFormat struct {
	children logrus.Formatter
}

func (t *tracingFormat) Format(e *logrus.Entry) ([]byte, error) {
	ctx := e.Context
	if ctx == nil {
		return t.children.Format(e)
	}

	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return t.children.Format(e)
	}

	data := e.Data
	jctx, ok := span.Context().(jaeger.SpanContext)
	if !ok {
		return t.children.Format(e)
	}
	data["trace-id"] = jctx.TraceID().String()
	e.Data = data
	return t.children.Format(e)
}

var (
	defaultLogger *logrus.Logger = NewLogger("info", nil)
)

func NewLogger(level string, output io.Writer) *logrus.Logger {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	log := logrus.New()
	log.SetLevel(logLevel)

	if output != nil {
		log.SetOutput(output)
	}

	log.SetFormatter(&callerFormat{
		children: &tracingFormat{
			children: &logrus.JSONFormatter{
				TimestampFormat: `2006-01-02 15:04:05`,
			},
		},
	})

	return log
}

func SetDefaultLogger(logger *logrus.Logger) {
	defaultLogger = logger
}

func WithContext(ctx context.Context) *logrus.Entry {
	return defaultLogger.WithContext(ctx)
}
