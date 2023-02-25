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
	"yehuizhang.com/go-webapp-gin/src/tasks"
)

var WireSet = wire.NewSet(wire.Struct(new(Server), "*"), RouterSet)

type Server struct {
	Router *Router
	Config *viper.Viper
	Redis  database.IRedis
	Pg     database.IPostgres
	Log    *logger.Logger
}

// InitGinEngine Graceful-shutdown : https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
func (s Server) InitGinEngine() (*gin.Engine, error) {

	if err := s.prepareDB(); err != nil {
		return nil, err
	}

	app := gin.New()
	s.Router.RegisterAPI(app)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	defer s.Redis.Client().Close()

	go func() {
		app.Run(s.Config.GetString("GIN_PORT"))
	}()

	// Listen for the interrupt signal
	<-ctx.Done()

	s.Redis.Client().Close()
	stop()
	s.Log.Info("shutting down gracefully, press Ctrl+C again to force")

	return app, nil
}

func (s Server) prepareDB() error {
	return tasks.AutoMigratePgSchema(s.Pg)
}
