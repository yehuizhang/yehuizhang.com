package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type IRedis interface {
	Client() *redis.Client
}

type Redis struct {
	client *redis.Client
}

func InitRedis(config *viper.Viper) (IRedis, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.GetString("REDIS_ADDRESS"),
		Password: config.GetString("REDIS_PASSWORD"),
		DB:       config.GetInt("REDIS_DEFAULT_DB"),
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("db: unable to establish connection to redis server. %s", err)

	}
	return &Redis{client: redisClient}, nil
}

func (r Redis) Client() *redis.Client {
	return r.client
}
