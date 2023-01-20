package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type Database struct {
	Redis *redis.Client
}

func InitDatabase(c *viper.Viper, log *logger.Logger) *Database {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.GetString("db.address"),
		Password: c.GetString("db.password"),
		DB:       c.GetInt("db.default_db"),
	})
	log.Info("DB initialization succeed")

	return &Database{Redis: redisClient}
}
