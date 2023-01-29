package main

import (
	"github.com/google/wire"
	"log"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/dao/user/account"
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
		log.Panic(err)
	}

	userAccountQuery := account.UserAccountQuery{DB: injector.Server.Database, Log: injector.Log}
	userAccountQuery.Create()

	injector.Server.InitGinEngine()
}
