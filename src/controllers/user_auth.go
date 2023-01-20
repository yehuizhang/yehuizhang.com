package controllers

import (
	"github.com/google/wire"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/logger"

	"github.com/gin-gonic/gin"
	"yehuizhang.com/go-webapp-gin/src/models/user"
	"yehuizhang.com/go-webapp-gin/src/utils/auth"
)

var UserAuthControllerSet = wire.NewSet(wire.Struct(new(UserAuthController), "*"))

type UserAuthController struct {
	logger      *logger.Logger
	infoHandler *user.InfoHandler
	authHandler *user.AuthHandler
}

func (ua *UserAuthController) SignUp(c *gin.Context) {

	userCredential, err := ua.authHandler.SignIn(c)
	if err != nil {
		if err == user.PasswordNotMatch {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Password does not match.",
			})
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	userInfo, err := ua.infoHandler.GetUserInfo(userCredential.ID)
	if err != nil {
		ua.logger.Errorw("Unable to get UserInfo.", "error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	auth.AddUidToSessionStore(c, userCredential.ID)
	c.JSON(http.StatusOK, userInfo)
}

func (ua *UserAuthController) Signup(c *gin.Context) {
	userInfo, userCredential, err := ua.authHandler.Signup(c)

	if err != nil {
		ua.logger.Errorw("signup failed: ", "error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	auth.AddUidToSessionStore(c, userCredential.ID)
	c.JSON(http.StatusCreated, *userInfo)
}
