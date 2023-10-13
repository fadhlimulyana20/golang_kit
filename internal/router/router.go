package router

import (
	"encoding/json"
	"net/http"
	"template/internal/appctx"
	"template/internal/config"
	m "template/internal/middleware"
	mail "template/utils/mailer"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type router struct {
	router *chi.Mux
	cfg    *config.Config
}

type RouterCfg struct {
	DB        *gorm.DB
	SMTP      mail.Mailer
	Secret    string
	AesSecret string
}

func NewRouter(cfg *config.Config) Router {
	return &router{
		router: chi.NewRouter(),
		cfg:    cfg,
	}
}

func (rtr *router) Route() http.Handler {
	rtr.router.Use(m.Cors(rtr.cfg.DB))
	rtr.router.Use(m.Logger)
	rtr.router.Use(m.Recovery)
	rtr.router.Use(m.Authorization(rtr.cfg.DB))
	rtr.router.Use(m.Pagination)

	rtr.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		logrus.Error("Error 404 page not found")
		resp := *appctx.NewResponse().WithErrors("Page not found").WithCode(404)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)
		d, _ := json.Marshal(resp)
		w.Write(d)
	})

	rtr.router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		logrus.Error("Error 405 method not allowed")
		resp := *appctx.NewResponse().WithErrors("method not allowed").WithCode(405)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)
		d, _ := json.Marshal(resp)
		w.Write(d)
	})

	rtr.router.Mount("/hello", rtr.helloRouter())
	rtr.router.Mount("/nosql", rtr.nosqlRouter())
	rtr.router.Mount("/admin/v1", rtr.AdminRouterV1())

	return rtr.router
}
