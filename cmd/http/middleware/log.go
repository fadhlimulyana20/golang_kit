package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.WithFields(log.Fields{
			"at":     time.Now().UTC(),
			"url":    request.URL.Path,
			"method": request.Method,
			"ip":     request.RemoteAddr,
		}).Info("Incoming Request")
		handler.ServeHTTP(writer, request)
	})
}
