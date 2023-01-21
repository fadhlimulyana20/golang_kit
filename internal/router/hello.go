package router

import (
	"net/http"
	"template/internal/handler"

	"github.com/go-chi/chi/v5"
)

func (rtr *router) helloRouter() http.Handler {
	helloHandler := handler.NewHelloHandler()
	hello := chi.NewRouter()
	hello.Get("/", helloHandler.Hello)
	return hello
}
