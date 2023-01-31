package info

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type IUserInfoQuery interface {
	Create(userInfo *UserInfo) int
	Get(id string) (*UserInfo, int)
}

type UserInfoQuery struct {
	Db  *database.Database
	Log *logger.Logger
}

func InitUserInfoQuery(db *database.Database, log *logger.Logger) IUserInfoQuery {
	return UserInfoQuery{
		Db:  db,
		Log: log,
	}
}

func (u UserInfoQuery) Create(userInfo *UserInfo) int {
	tx := u.Db.Pg.Create(userInfo)
	if tx.Error != nil {
		u.Log.Errorw("failed to store Userinfo into DB.", "err", tx.Error)
		return http.StatusInternalServerError
	}
	return 0
}

func (u UserInfoQuery) Get(id string) (*UserInfo, int) {
	var result UserInfo
	tx := u.Db.Pg.First(&result, "id = ?", id)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		u.Log.Errorw("userinfo was not found", "id", id)
		return nil, http.StatusNotFound
	}
	if tx.Error != nil {
		u.Log.Errorw("failed to retrieve userInfo from DB", "id", id, "err", tx.Error)
		return nil, http.StatusInternalServerError
	}
	return &result, 0
}
