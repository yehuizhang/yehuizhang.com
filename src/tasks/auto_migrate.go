package tasks

import (
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/src/dao/user/account"
)

func AutoMigratePgSchema(db *database.Database) error {

	err := db.Pg.AutoMigrate(&account.UserAccount{})
	return err
}
