package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/forms"
	"yehuizhang.com/go-webapp-gin/src/models/user"
)

var UserInfoControllerSet = wire.NewSet(wire.Struct(new(UserInfoController), "*"))

type UserInfoController struct {
	logger      *logger.Logger
	infoHandler *user.InfoHandler
}

func (ui *UserInfoController) Get(c *gin.Context) {
	uid := c.GetString(user.UID)
	userInfo, err := ui.infoHandler.GetUserInfo(uid)
	if err != nil {
		ui.logger.Errorw("Unable to get UserInfo.", "error: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

func (ui *UserInfoController) CreateOrUpdate(c *gin.Context) {
	uid := c.GetString(user.UID)

	userInfoInput := forms.UserInfo{}
	err := c.Bind(&userInfoInput)

	if err != nil {
		ui.logger.Errorw("Unable to read UserInfo from body: ", "error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var userInfo *user.UserInfo
	userInfo, err = ui.infoHandler.AddOrUpdate(uid, userInfoInput)

	if err != nil {
		ui.logger.Errorw("Unable to update UserInfo", "error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, *userInfo)
}
