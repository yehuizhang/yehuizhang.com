package server

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"yehuizhang.com/go-webapp-gin/src/config"
	"yehuizhang.com/go-webapp-gin/src/database"
)

type Server interface{}

// Graceful-shutdown : https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
func NewServer(cfg config.Config, db *database.Database) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	defer db.Redis.Close()

	r := NewRouter(db)

	go func() {
		r.Run(cfg.GetString("server.port"))
	}()
	// Listen for the interrupt signal
	<-ctx.Done()

	db.Redis.Close()
	stop()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

}
