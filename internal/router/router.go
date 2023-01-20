package router

import (
	"encoding/json"
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

func (r *router) Route() http.Handler {
	router := chi.NewRouter()
	router.Use(m.Logger)
	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		re := map[string]string{
			"hello": "world",
		}
		res, _ := json.Marshal(re)
		w.Write(res)
	})

	return router
}
