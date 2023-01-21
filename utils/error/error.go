package error

import (
	"errors"
	"fmt"
)

type ErrorWithContext struct {
	err error
	ctx interface{}
}

func NewWithContext(ctx interface{}, msg string) ErrorWithContextInf {
	return &ErrorWithContext{
		err: errors.New(msg),
		ctx: ctx,
	}
}

func (e *ErrorWithContext) Error() string {
	return fmt.Sprintf("%s. Context: %v", e.err.Error(), e.ctx)
}
