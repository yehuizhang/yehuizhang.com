package account

import (
	"github.com/pkg/errors"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

type UserAccountQuery struct {
	DB  *database.Database
	Log *logger.Logger
}

func (u UserAccountQuery) Create() error {

	u.Log.Info("Creating new user: Starting")
	user := UserAccount{
		Username: "yehuizhang",
		Password: "password",
		Email:    "yehuizhang@test.com",
		Active:   true,
	}
	result := u.DB.Pg.Model(new(UserAccount)).Create(&user)
	u.Log.Info("Creating new user: finished")

	return errors.WithStack(result.Error)

}
