package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func startServer() {
	log.Info("HTTP server running in port 3000")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}

var StartServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	Long:  "Start HTTP Server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}
