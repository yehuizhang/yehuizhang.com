package flag_parser

import (
	"flag"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type FlagParser struct {
	Env        string
	ConfigPath string
	ConfigName string
	ConfigType string
}

func InitFlagParser(log *logger.Logger) *FlagParser {
	env := flag.String("env", "local", "environment: {local|test|development|production}")
	configPath := flag.String("configPath", ".", "path of the configuration file")
	configName := flag.String("configName", ".env", "name of the configuration file")
	configType := flag.String("configType", "env", "type of the configuration file")
	flag.Parse()
	log.Infof("Environment: %s", *env)
	return &FlagParser{Env: *env, ConfigPath: *configPath, ConfigName: *configName, ConfigType: *configType}
}
