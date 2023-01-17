package database

import (
	"log"

	"github.com/go-redis/redis/v8"
	"yehuizhang.com/go-webapp-gin/src/config"
)

type Database struct {
	Redis *redis.Client
}

func InitDatabase(c config.Config) *Database {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.GetString("db.address"),
		Password: c.GetString("db.password"),
		DB:       c.GetInt("db.default_db"),
	})
	log.Println("Status: DB initialization succeed")

	return &Database{Redis: redisClient}
}
