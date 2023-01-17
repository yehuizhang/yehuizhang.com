package config

import (
	"fmt"
	"yehuizhang.com/go-webapp-gin/pkg/logger"

	"github.com/spf13/viper"
)

type Config = *viper.Viper

func InitConfig(flagParser *FlagParser, logger *logger.Logger) (Config, error) {

	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(flagParser.env)
	config.AddConfigPath("../../config/")
	config.AddConfigPath("config/")
	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error on parsing configuration file")
	}
	logger.Info("Status: Config initialization succeed")

	return config, nil
}
