package flag_parser

import (
	"flag"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type FlagParser struct {
	Env string
}

func InitFlagParser(log *logger.Logger) *FlagParser {
	env := flag.String("env", "local", "environment: {local|test|development|production}")
	flag.Parse()
	log.Infof("Environment: %s", *env)
	return &FlagParser{Env: *env}
}
