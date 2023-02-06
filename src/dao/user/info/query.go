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
	Update(userInfo *UserInfo) int
}

type UserInfoQuery struct {
	Pg  database.IPostgres
	Log *logger.Logger
}

func InitUserInfoQuery(pg database.IPostgres, log *logger.Logger) IUserInfoQuery {
	return UserInfoQuery{
		Pg:  pg,
		Log: log,
	}
}

func (u UserInfoQuery) Create(userInfo *UserInfo) int {
	tx := u.Pg.Client().Create(userInfo)
	if tx.Error != nil {
		u.Log.Errorw("failed to store Userinfo into DB.", "err", tx.Error)
		return http.StatusInternalServerError
	}
	return 0
}

func (u UserInfoQuery) Update(userInfo *UserInfo) int {
	tx := u.Pg.Client().Save(userInfo)
	if tx.Error != nil {
		u.Log.Error("failed to update Userinfo into DB.", tx.Error)
		return http.StatusInternalServerError
	}
	return 0
}

func (u UserInfoQuery) Get(id string) (*UserInfo, int) {
	var result UserInfo
	tx := u.Pg.Client().First(&result, "id = ?", id)

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
