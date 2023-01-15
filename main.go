package main

import (
	"flag"
	"fmt"
	"os"

	"yehuizhang.com/go-webapp-gin/config"
	"yehuizhang.com/go-webapp-gin/db"
	"yehuizhang.com/go-webapp-gin/server"
)

func main() {
	environment := flag.String("e", "local", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {local|test|development|production}")
		os.Exit(1)
	}
	flag.Parse()

	fmt.Printf("Environment: %s\n", *environment)
	config.Init(*environment)

	db.Init()
	server.Init()
}
