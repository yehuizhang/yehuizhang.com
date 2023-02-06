package tasks

import (
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/src/dao/user/account"
	"yehuizhang.com/go-webapp-gin/src/dao/user/info"
)

func AutoMigratePgSchema(pg database.IPostgres) error {

	err := pg.Client().AutoMigrate(&account.UserAccount{}, &info.UserInfo{})
	return err
}
