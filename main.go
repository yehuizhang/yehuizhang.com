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

//	@title			zyh-go-webserver
//	@version		1.0
//	@description	This is the backend server for yehuizhang.com
//	@termsOfService	http://yehuizhang.com/terms/

//	@contact.name	Yehui Zhang
//	@contact.url	http://www.yehuizhang.com/support
//	@contact.email	yehuizhang@yehuizhang.com

//	@license.name	MIT License
//	@license.url	https://github.com/yehuizhang/go-zyh-webserver/blob/main/LICENSE

//	@host		localhost:8080
//	@BasePath	/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	injector, _, err := BuildInjector()

	if err != nil {
		log.Panic(err)
	}

	injector.Server.InitGinEngine()
}
