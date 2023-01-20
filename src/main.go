package main

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/server"
)

type Injector struct {
	server *server.Server
	log    *logger.Logger
}

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

func main() {
	injector, _, err := BuildInjector()

	if err != nil {
		injector.log.Error(err)
	}
	injector.server.InitGinEngine()
}
