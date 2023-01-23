package main

import (
	"github.com/google/wire"
	"log"
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

	//userAccountQuery := user_account.UserAccountQuery{DB: injector.Server.Database}
	//userAccountQuery.Create()

	if err != nil {
		log.Panic(err)
	}
	injector.Server.InitGinEngine()
}
