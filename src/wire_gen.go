// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/flag_parser"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/config"
	"yehuizhang.com/go-webapp-gin/src/controllers"
	"yehuizhang.com/go-webapp-gin/src/models/user"
	"yehuizhang.com/go-webapp-gin/src/server"
)

// Injectors from wire.go:

func BuildInjector() (Injector, func(), error) {
	healthController := &controllers.HealthController{}
	sugaredLogger := logger.InitLogger()
	flagParser := flag_parser.InitFlagParser(sugaredLogger)
	viper, err := config.InitConfig(flagParser, sugaredLogger)
	if err != nil {
		return Injector{}, nil, err
	}
	databaseDatabase, err := database.InitDatabase(viper, sugaredLogger)
	if err != nil {
		return Injector{}, nil, err
	}
	infoHandler := &user.InfoHandler{
		Database: databaseDatabase,
		Log:      sugaredLogger,
	}
	authHandler := &user.AuthHandler{
		Database: databaseDatabase,
		Log:      sugaredLogger,
	}
	userAuthController := &controllers.UserAuthController{
		Logger:      sugaredLogger,
		InfoHandler: infoHandler,
		AuthHandler: authHandler,
	}
	userInfoController := &controllers.UserInfoController{
		Logger:      sugaredLogger,
		InfoHandler: infoHandler,
	}
	router := &server.Router{
		HealthController:   healthController,
		UserAuthController: userAuthController,
		UserInfoController: userInfoController,
		Database:           databaseDatabase,
	}
	serverServer := &server.Server{
		Router:   router,
		Config:   viper,
		Database: databaseDatabase,
		Log:      sugaredLogger,
	}
	injector := Injector{
		Server: serverServer,
		Log:    sugaredLogger,
	}
	return injector, func() {
	}, nil
}
