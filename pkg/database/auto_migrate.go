package database

import "yehuizhang.com/go-webapp-gin/src/dao/pg/user_account"

func AutoMigratePgSchema(db *Database) error {

	err := db.Pg.AutoMigrate(&user_account.UserAccount{})
	return err
}
