package config

import (
	"fmt"
	"yehuizhang.com/go-webapp-gin/pkg/flag_parser"
	"yehuizhang.com/go-webapp-gin/pkg/logger"

	"github.com/spf13/viper"
)

func InitConfig(flagParser *flag_parser.FlagParser, logger *logger.Logger) (*viper.Viper, error) {

	config := viper.New()
	config.SetConfigType(flagParser.ConfigType)
	config.SetConfigName(flagParser.ConfigName)
	config.AddConfigPath(flagParser.ConfigPath)
	config.AddConfigPath(".")
	config.AddConfigPath("../")
	config.AddConfigPath("../../")
	if err := config.ReadInConfig(); err != nil {
		logger.Info("Config initialization: FAILED")
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		}
		return nil, fmt.Errorf("error occurred when initializing config. %s", err)
	}

	logger.Info("Config initialization: OK")

	return config, nil
}
