package flag_parser

import (
	"flag"
	"fmt"
)

type FlagParser struct {
	Env string
}

func InitFlagParser() *FlagParser {
	env := flag.String("env", "local", "environment: {local|test|development|production}")
	flag.Parse()
	fmt.Printf("Environment: %s\n", *env)

	return &FlagParser{Env: *env}
}
