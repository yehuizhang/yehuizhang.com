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
	Router   *Router
	Config   *viper.Viper
	Database *database.Database
	Log      *logger.Logger
}

// InitGinEngine Graceful-shutdown : https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
func (s Server) InitGinEngine() *gin.Engine {

	app := gin.New()
	s.Router.RegisterAPI(app)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	defer s.Database.Redis.Close()

	go func() {
		app.Run(s.Config.GetString("GIN_PORT"))
	}()

	// Listen for the interrupt signal
	<-ctx.Done()

	s.Database.Redis.Close()
	stop()
	s.Log.Info("shutting down gracefully, press Ctrl+C again to force")

	return app
}
