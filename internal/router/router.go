package router

import (
	"net/http"
	m "template/internal/middleware"

	"github.com/go-chi/chi/v5"
)

type router struct {
	router *chi.Mux
}

func NewRouter() Router {
	return &router{
		router: chi.NewRouter(),
	}
}

func (rtr *router) Route() http.Handler {
	rtr.router.Use(m.Logger)
	rtr.router.Use(m.Authorization)
	rtr.router.Use(m.Pagination)

	rtr.router.Mount("/hello", rtr.helloRouter())
	rtr.router.Mount("/user", rtr.UserRouter())

	return rtr.router
}
