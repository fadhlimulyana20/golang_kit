package error

import (
	"errors"
	"fmt"
	"strings"
)

type ErrorWithContext struct {
	Err error
	Ctx map[string]interface{}
}

func NewWithContext(ctx map[string]interface{}, msg string) ErrorWithContextInf {
	return &ErrorWithContext{
		Err: errors.New(msg),
		Ctx: ctx,
	}
}

func (e *ErrorWithContext) Error() string {
	var ctx []string
	for k, v := range e.Ctx {
		s := fmt.Sprintf("key: %s, msg: %s", k, v)
		ctx = append(ctx, s)
	}
	return fmt.Sprintf("%s. Context: [%s]", e.Err.Error(), strings.Join(ctx, ","))
}
