//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/flag_parser"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/config"
	"yehuizhang.com/go-webapp-gin/src/models"
	"yehuizhang.com/go-webapp-gin/src/server"
)

func BuildInjector() (Injector, func(), error) {
	wire.Build(
		logger.InitLogger,
		flag_parser.InitFlagParser,
		config.InitConfig,
		database.InitDatabase,
		models.ModelsSet,
		server.ServerSet,
		InjectorSet,
	)

	return Injector{}, nil, nil
}
