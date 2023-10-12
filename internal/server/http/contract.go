package http

import (
	"context"
)

type Server interface {
	Run(ctx context.Context, port int)
	Done()
}
