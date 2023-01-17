//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/config"
	"yehuizhang.com/go-webapp-gin/src/controllers"
	"yehuizhang.com/go-webapp-gin/src/database"
	"yehuizhang.com/go-webapp-gin/src/server"
)

func BuildInjector() (Injector, func(), error) {
	wire.Build(
		logger.InitLogger,
		config.InitFlagParser,
		config.InitConfig,
		controllers.ControllerSet,
		server.ServerSet,
		database.InitDatabase,
		InjectorSet,
	)

	return Injector{}, nil, nil
}
