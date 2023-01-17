package main

import (
	"flag"
	"fmt"

	"yehuizhang.com/go-webapp-gin/src/config"
	"yehuizhang.com/go-webapp-gin/src/database"
	"yehuizhang.com/go-webapp-gin/src/server"
)

func main() {

	env := parseFlag()
	c := config.NewConfig(env)
	db := database.NewDatabase(c)
	server.NewServer(c, db)
}

func parseFlag() string {

	// flag.Usage = func() {
	// 	fmt.Println("Usage: server -env {local|test|development|production}")
	// 	os.Exit(1)
	// }
	environment := flag.String("env", "local", "environment: {local|test|development|production}")
	flag.Parse()
	fmt.Printf("Environment: %s\n", *environment)
	return *environment
}
