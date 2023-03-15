package router

import (
	"net/http"
	m "template/internal/middleware"
	mail "template/utils/mailer"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type router struct {
	router *chi.Mux
}

type RouterCfg struct {
	DB        *gorm.DB
	SMTP      mail.Mailer
	Secret    string
	AesSecret string
}

func NewRouter(r *RouterCfg) Router {
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
