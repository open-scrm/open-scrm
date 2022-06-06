package vo

import (
	"context"
	"fmt"
)

type Error struct {
	msg string
}

func NewError(ctx context.Context, format string, args ...interface{}) error {
	return &Error{
		msg: fmt.Sprintf(format, args...),
	}
}

func (e Error) Error() string {
	return e.msg
}
