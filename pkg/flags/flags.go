package flags

import (
	"flag"
	"log"
)

type FlagParser struct {
	Env        string
	ConfigPath string
	ConfigName string
	ConfigType string
}

func InitFlagParser() *FlagParser {
	env := flag.String("env", "local", "environment: {local|test|dev|prod}")
	configPath := flag.String("configPath", ".", "path of the configuration file")
	configName := flag.String("configName", ".env", "name of the configuration file")
	configType := flag.String("configType", "env", "type of the configuration file")
	flag.Parse()
	f := &FlagParser{Env: *env, ConfigPath: *configPath, ConfigName: *configName, ConfigType: *configType}
	log.Printf("flags: %+v", *f)
	return f
}
