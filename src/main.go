package main

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/server"
)

type Injector struct {
	Server *server.Server
	Log    *logger.Logger
}

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

func main() {
	injector, _, err := BuildInjector()

	if err != nil {
		injector.Log.Error(err)
	}
	injector.Log.Info("succeed")
	injector.Server.InitGinEngine()
}
