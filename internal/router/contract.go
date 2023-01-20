package router

import "net/http"

type Router interface {
	Route() http.Handler
}
