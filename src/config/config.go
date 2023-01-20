package config

import (
	"fmt"
	"yehuizhang.com/go-webapp-gin/pkg/flag_parser"
	"yehuizhang.com/go-webapp-gin/pkg/logger"

	"github.com/spf13/viper"
)

func InitConfig(flagParser *flag_parser.FlagParser, logger *logger.Logger) (*viper.Viper, error) {

	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(flagParser.Env)
	config.AddConfigPath("../../config/")
	config.AddConfigPath("config/")
	if err := config.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error on parsing configuration file")
	}
	logger.Info("Status: Config initialization succeed")

	return config, nil
}
