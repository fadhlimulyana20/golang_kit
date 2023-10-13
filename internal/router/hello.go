package router

import (
	"net/http"
	"template/internal/handler"

	"github.com/go-chi/chi/v5"
)

func (rtr *router) helloRouter() http.Handler {
	helloHandler := handler.NewHelloHandler()
	uploadHandler := handler.NewUploadHandler()
	hello := chi.NewRouter()
	hello.Get("/", helloHandler.Hello)
	hello.Post("/upload", uploadHandler.Upload)
	return hello
}

func (rtr *router) nosqlRouter() http.Handler {
	nosqlHandler := handler.NewNoSQLHandler(rtr.cfg.Mongo)
	nosql := chi.NewRouter()
	nosql.Post("/", nosqlHandler.Insert)
	return nosql
}
