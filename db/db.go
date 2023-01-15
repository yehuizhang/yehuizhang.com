package db

import (
	"log"

	"github.com/go-redis/redis/v8"
	"yehuizhang.com/go-webapp-gin/config"
)

var db_redis *redis.Client

func Init() {

	c := config.GetConfig()

	db_redis = redis.NewClient(&redis.Options{
		Addr:     c.GetString("db.address"),
		Password: c.GetString("db.password"),
		DB:       c.GetInt("db.default_db"),
	})
	log.Println("Status: DB initialization succeed")
}

func GetRedisDB() *redis.Client {
	return db_redis
}
