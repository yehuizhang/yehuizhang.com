package server

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"yehuizhang.com/go-webapp-gin/config"
	"yehuizhang.com/go-webapp-gin/db"
)

// Graceful-shutdown : https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
func Init() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	defer db.GetRedisDB().Close()

	config := config.GetConfig()
	r := NewRouter()

	go func() {
		r.Run(config.GetString("server.port"))
	}()
	// Listen for the interrupt signal
	<-ctx.Done()

	db.GetRedisDB().Close()
	stop()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

}
