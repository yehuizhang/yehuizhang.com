package user_account

import (
	"github.com/pkg/errors"
	"yehuizhang.com/go-webapp-gin/pkg/database"
)

type UserAccountQuery struct {
	DB *database.Database
}

func (u UserAccountQuery) Create() error {
	user := UserAccount{
		Username: "yehuizhang",
		Password: "password",
		Email:    "yehuizhang@test.com",
		Active:   true,
	}
	result := u.DB.Pg.Model(new(UserAccount)).Create(&user)

	return errors.WithStack(result.Error)

}
