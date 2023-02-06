// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/account"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/flags"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/config"
	"yehuizhang.com/go-webapp-gin/src/controllers/admin"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
	"yehuizhang.com/go-webapp-gin/src/server"
)

// Injectors from wire.go:

func BuildInjector() (Injector, func(), error) {
	controller := &admin.Controller{}
	flagParser := flags.InitFlagParser()
	sugaredLogger := logger.InitLogger(flagParser)
	viper, err := config.InitConfig(flagParser, sugaredLogger)
	if err != nil {
		return Injector{}, nil, err
	}
	iPostgres, err := database.InitPostgres(viper)
	if err != nil {
		return Injector{}, nil, err
	}
	iUserAccountQuery := account.InitUserAccountQuery(iPostgres, sugaredLogger)
	iUserInfoQuery := info.InitUserInfoQuery(iPostgres, sugaredLogger)
	userController := &user.Controller{
		Log:          sugaredLogger,
		AccountQuery: iUserAccountQuery,
		InfoQuery:    iUserInfoQuery,
	}
	iRedis, err := database.InitRedis(viper)
	if err != nil {
		return Injector{}, nil, err
	}
	router := &server.Router{
		AdminController: controller,
		UserController:  userController,
		Redis:           iRedis,
		Log:             sugaredLogger,
	}
	serverServer := &server.Server{
		Router: router,
		Config: viper,
		Redis:  iRedis,
		Pg:     iPostgres,
		Log:    sugaredLogger,
	}
	injector := Injector{
		Server: serverServer,
		Log:    sugaredLogger,
	}
	return injector, func() {
	}, nil
}
