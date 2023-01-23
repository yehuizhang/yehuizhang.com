package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type Database struct {
	Redis *redis.Client
	Pg    *gorm.DB
}

func InitDatabase(config *viper.Viper, log *logger.Logger) (*Database, error) {

	redisClient, err := initRedis(config)

	if err != nil {
		log.Info("DB initialization: FAILED")
		return nil, err
	}

	pgClient, err := initPostgres(config)

	if err != nil {
		log.Info("DB initialization: FAILED")
		return nil, err
	}

	log.Info("DB initialization: OK")

	return &Database{Redis: redisClient, Pg: pgClient}, nil
}

func initRedis(config *viper.Viper) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.GetString("REDIS_ADDRESS"),
		Password: config.GetString("REDIS_PASSWORD"),
		DB:       config.GetInt("REDIS_DEFAULT_DB"),
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, generateErrorMessage("redis", err.Error())
	}

	return redisClient, nil

}

func initPostgres(config *viper.Viper) (*gorm.DB, error) {
	host := config.GetString("PG_HOST")
	user := config.GetString("PG_USER")
	password := config.GetString("PG_PASSWORD")
	dbName := config.GetString("PG_DB_NAME")
	port := config.GetString("PG_PORT")
	sslMode := config.GetString("PG_SSL_MODE")
	timezone := config.GetString("PG_TIMEZONE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbName, port, sslMode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, generateErrorMessage("postgres", err.Error())
	}
	return db, nil
}

func generateErrorMessage(name string, errMsg string) error {
	return fmt.Errorf("db: unable to establish connection to %s server. %s", name, errMsg)
}
