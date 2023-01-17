package config

import (
	"flag"
	"fmt"
)

type FlagParser struct {
	env string
}

func InitFlagParser() *FlagParser {
	env := flag.String("env", "local", "environment: {local|test|development|production}")
	flag.Parse()
	fmt.Printf("Environment: %s\n", *env)

	return &FlagParser{env: *env}
}
