package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IPostgres interface {
	Client() *gorm.DB
}

type Postgres struct {
	client *gorm.DB
}

func InitPostgres(config *viper.Viper) (IPostgres, error) {

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
		return nil, fmt.Errorf("db: unable to establish connection to redis server. %s", err)
	}
	return &Postgres{client: db}, nil
}

func (p Postgres) Client() *gorm.DB {
	return p.client
}
