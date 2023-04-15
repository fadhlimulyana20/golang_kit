package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func Cors(db *gorm.DB) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		cfgAllowedOrigin := os.Getenv("ALLOWED_HOST")

		allowedOrigin := []string{"https://*", "http://*"}

		if cfgAllowedOrigin != "*" {
			allowedOrigin = strings.Split(cfgAllowedOrigin, ",")
		}

		h := cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: allowedOrigin,
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})

		return h(handler)
	}
}
