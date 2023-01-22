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
	Logger      *logger.Logger
	InfoHandler *user.InfoHandler
	AuthHandler *user.AuthHandler
}

func (ua *UserAuthController) SignIn(c *gin.Context) {

	userCredential, err := ua.AuthHandler.SignIn(c)
	if err != nil {
		if err == user.PasswordNotMatch {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Password does not match.",
			})
		} else {
			ua.Logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	userInfo, err := ua.InfoHandler.GetUserInfo(userCredential.ID)
	if err != nil {
		ua.Logger.Errorw("Unable to get UserInfo.", "error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	auth.AddUidToSessionStore(c, userCredential.ID)
	c.JSON(http.StatusOK, userInfo)
}

func (ua *UserAuthController) SignUp(c *gin.Context) {
	userInfo, userCredential, err := ua.AuthHandler.Signup(c)

	if err != nil {
		ua.Logger.Errorw("signup failed: ", "error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	auth.AddUidToSessionStore(c, userCredential.ID)
	c.JSON(http.StatusCreated, *userInfo)
}
