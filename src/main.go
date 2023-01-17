package main

import (
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/src/server"
)

type Injector struct {
	server *server.Server
}

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

func main() {

}
