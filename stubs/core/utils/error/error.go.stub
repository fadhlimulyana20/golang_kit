package error

import (
	"errors"
	"fmt"
	"strings"
)

type ErrorWithContext struct {
	err error
	ctx []string
}

func NewWithContext(ctx []string, msg string) ErrorWithContextInf {
	return &ErrorWithContext{
		err: errors.New(msg),
		ctx: ctx,
	}
}

func (e *ErrorWithContext) Error() string {
	return fmt.Sprintf("%s. Context: [%s]", e.err.Error(), strings.Join(e.ctx, ","))
}
