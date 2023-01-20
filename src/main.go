package main

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type Injector struct {
	//server *server.Server
	logger *logger.Logger
	config *viper.Viper
}

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

func main() {

}
