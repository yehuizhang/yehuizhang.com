//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/flag_parser"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/config"
)

func BuildInjector() (Injector, func(), error) {
	wire.Build(
		logger.InitLogger,
		flag_parser.InitFlagParser,
		config.InitConfig,
		//controllers.ControllerSet,
		//server.ServerSet,
		//database.InitDatabase,
		InjectorSet,
	)

	return Injector{}, nil, nil
}
