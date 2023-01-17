package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config = *viper.Viper

func NewConfig(env string) Config {

	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../../configs/")
	config.AddConfigPath("configs/")
	if err := config.ReadInConfig(); err != nil {
		log.Panic("error on parsing configuration file")
	}
	log.Println("Status: Config initialization succeed")

	return config
}

// func relativePath(basedir string, path *string) {
// 	p := *path
// 	if len(p) > 0 && p[0] != '/' {
// 		*path = filepath.Join(basedir, p)
// 	}
// }
