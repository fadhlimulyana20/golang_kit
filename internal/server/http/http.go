package http

import (
	"context"
	"fmt"
	"net/http"
	"template/internal/config"
	"template/internal/router"
	"template/utils/mailer"
	"template/utils/minio"

	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type httpServer struct {
	router http.Handler
}

func NewServer(env string, db *gorm.DB, smtp mailer.Mailer, secretKey string, minio minio.MinioStorageContract) Server {
	// Init Config
	cfg := &config.Config{
		ENV:    env,
		DB:     db,
		SMTP:   &smtp,
		Secret: secretKey,
		Minio:  minio,
	}

	// Create new router
	rtr := router.NewRouter(cfg)
	return &httpServer{
		router: rtr.Route(),
	}
}

// Run server
func (h *httpServer) Run(ctx context.Context, port int) {
	log.Info(fmt.Sprintf("Server running on port %d. Access it from http://127.0.0.1:%d\n", port, port))
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: h.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Warn("http server got %w", err.Error())
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Info("server existed properly")
}

func (h *httpServer) Done() {
	log.Fatal("Server closed")
}
