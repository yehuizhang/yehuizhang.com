package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type Database struct {
	Redis *redis.Client
}

func InitDatabase(c *viper.Viper, log *logger.Logger) (*Database, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.GetString("db.address"),
		Password: c.GetString("db.password"),
		DB:       c.GetInt("db.default_db"),
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Info("DB initialization: FAILED")
		return nil, fmt.Errorf("db: unable to establish connection to the server. %s", err)
	}

	log.Info("DB initialization: OK")

	return &Database{Redis: redisClient}, nil
}
