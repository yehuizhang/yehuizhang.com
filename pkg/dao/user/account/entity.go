package account

import (
	"yehuizhang.com/go-webapp-gin/pkg/dao/shared"
)

type UserAccount struct {
	shared.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Active   bool   `gorm:"not null"`
}

type SignUpForm struct {
	// This should be improved by using custom validator
	Username string `json:"username" binding:"required,alphanum,lowercase,min=3,max=15"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Email    string `json:"email" binding:"required,email"`
}

type SignInForm struct {
	Username string `json:"username" binding:"required,alphanum,lowercase,min=3,max=15"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}
