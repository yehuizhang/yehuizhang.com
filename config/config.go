package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {

	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	log.Println("Status: Config initialization succeed")
}

func GetConfig() *viper.Viper {
	return config
}

// func relativePath(basedir string, path *string) {
// 	p := *path
// 	if len(p) > 0 && p[0] != '/' {
// 		*path = filepath.Join(basedir, p)
// 	}
// }
