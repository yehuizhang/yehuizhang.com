package tasks

import (
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/src/dao/pg/user_account"
)

func AutoMigratePgSchema(db *database.Database) error {

	err := db.Pg.AutoMigrate(&user_account.UserAccount{})
	return err
}
