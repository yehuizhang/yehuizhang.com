package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"os/signal"
	"syscall"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

var ServerSet = wire.NewSet(wire.Struct(new(Server), "*"), RouterSet)

type Server struct {
	router   *Router
	config   *viper.Viper
	database *database.Database
	log      *logger.Logger
}

// InitGinEngine Graceful-shutdown : https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
func (s Server) InitGinEngine() *gin.Engine {

	app := gin.New()
	s.router.RegisterAPI(app)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	defer s.database.Redis.Close()

	go func() {
		app.Run(s.config.GetString("server.port"))
	}()

	// Listen for the interrupt signal
	<-ctx.Done()

	s.database.Redis.Close()
	stop()
	s.log.Info("shutting down gracefully, press Ctrl+C again to force")

	return app
}
