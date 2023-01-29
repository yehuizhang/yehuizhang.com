package account

import (
	"yehuizhang.com/go-webapp-gin/src/dao/shared"
)

type UserAccount struct {
	shared.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Active   bool   `gorm:"not null"`
}
