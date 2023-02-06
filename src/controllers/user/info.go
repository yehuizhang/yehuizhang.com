package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
)

func (ctl *Controller) GetInfo(c *gin.Context) {
	uid := c.GetString(UID)

	userInfo, errorCode := ctl.InfoQuery.Get(uid)

	if errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

func (ctl *Controller) CreateInfo(c *gin.Context) {
	uid := c.GetString(UID)
	input := info.Form{}
	if err := c.ShouldBind(&input); err != nil {
		ctl.Log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userInfo := info.UserInfo{
		Id:       uid,
		Name:     input.Name,
		Birthday: input.Birthday,
		Gender:   input.Gender,
		PhotoURL: input.PhotoURL,
	}

	if errorCode := ctl.InfoQuery.Create(&userInfo); errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

func (ctl *Controller) UpdateInfo(c *gin.Context) {
	uid := c.GetString(UID)
	input := info.Form{}
	if err := c.ShouldBind(&input); err != nil {
		ctl.Log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userInfo, errorCode := ctl.InfoQuery.Get(uid)
	if errorCode != 0 {
		ctl.Log.Errorf("unable to get userInfo from db for user %s", uid)
		c.AbortWithStatus(errorCode)
		return
	}

	userInfo.Name = input.Name
	userInfo.Birthday = input.Birthday
	userInfo.Gender = input.Gender
	userInfo.PhotoURL = input.PhotoURL

	if errorCode := ctl.InfoQuery.Update(userInfo); errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
