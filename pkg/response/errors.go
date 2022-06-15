package response

import (
	"context"
	"fmt"
)

type Error struct {
	code  int
	msg   string
	cause error
}

func NewError(ctx context.Context, format string, args ...interface{}) error {
	return &Error{
		msg:  fmt.Sprintf(format, args...),
		code: -1,
	}
}

func NewErrorCode(code int, format string, args ...interface{}) error {
	return &Error{
		msg:  fmt.Sprintf(format, args...),
		code: code,
	}
}

func InvalidParam(err error) error {
	return &Error{
		msg:   "参数错误",
		code:  -1,
		cause: err,
	}
}

func (e Error) Error() string {
	return e.msg
}
