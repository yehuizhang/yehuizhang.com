package account

import (
	"errors"
	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

var WireSet = wire.NewSet(wire.Struct(new(UserAccountQuery), "*"))

type UserAccountQuery struct {
	Db  *database.Database
	Log *logger.Logger
}

func (u UserAccountQuery) Create(input *SignUpForm) (string, int) {

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 4)
	if err != nil {
		u.Log.Errorf("failed to encrypt user's password. %s", err)
		return "", http.StatusInternalServerError
	}
	userAccount := UserAccount{
		Username: input.Username,
		Password: string(encryptedPassword),
		Email:    input.Email,
		Active:   true,
	}
	tx := u.Db.Pg.Create(&userAccount)

	if tx.Error != nil {
		u.Log.Errorf("failed to store user account in DB. %s", tx.Error)
		return "", http.StatusInternalServerError
	}
	u.Log.Debugf("user %s is created", userAccount.Uuid.String())
	return userAccount.Uuid.String(), 0
}

func (u UserAccountQuery) GetByUsername(username string) (*UserAccount, int) {
	var record UserAccount

	tx := u.Db.Pg.Where("username = ?", username).First(&record)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		u.Log.Errorf("user %s was not found", username)
		return nil, http.StatusNotFound
	}

	if tx.Error != nil {
		u.Log.Errorw("failed to retrieve account record.", "err", tx.Error)
		return nil, http.StatusInternalServerError
	}
	return &record, 0
}
