//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/flags"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/config"
	"yehuizhang.com/go-webapp-gin/src/dao"
	"yehuizhang.com/go-webapp-gin/src/server"
)

func BuildInjector() (Injector, func(), error) {
	wire.Build(
		flags.InitFlagParser,
		logger.InitLogger,
		config.InitConfig,
		database.InitDatabase,
		dao.QuerySet,
		server.WireSet,
		InjectorSet,
	)

	return Injector{}, nil, nil
}
