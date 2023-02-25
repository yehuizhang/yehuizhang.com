package tasks

import (
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/account"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
	"yehuizhang.com/go-webapp-gin/pkg/database"
)

func AutoMigratePgSchema(pg database.IPostgres) error {

	err := pg.Client().AutoMigrate(&account.UserAccount{}, &info.UserInfo{})
	return err
}
