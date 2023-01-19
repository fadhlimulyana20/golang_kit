package http

import (
	"encoding/json"
	"net/http"
	m "template/cmd/http/middleware"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func startServer() {
	log.Info("HTTP server running in port 3000")
	r := chi.NewRouter()
	r.Use(m.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		re := map[string]string{
			"hello": "world",
		}
		res, _ := json.Marshal(re)
		w.Write(res)
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
