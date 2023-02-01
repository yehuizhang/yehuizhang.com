package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"yehuizhang.com/go-webapp-gin/src/dao/user/info"
)

func (ctl *Controller) Get(c *gin.Context) {
	uid := c.GetString(UID)

	userInfo, errorCode := ctl.InfoQuery.Get(uid)

	if errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

func (ctl *Controller) Create(c *gin.Context) {
	uid := c.GetString(UID)
	input := info.Form{}
	err := c.Bind(&input)

	if err != nil {
		ctl.Log.Error("Unable to read UserInfo from body.", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var birthday time.Time
	if input.Birthday != "" {
		birthday, err = time.Parse("2006-01-02", input.Birthday)
		if err != nil {
			ctl.Log.Error("Unable to parse birthday from input", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	userInfo := info.UserInfo{
		Uuid:     uid,
		Name:     input.Name,
		Birthday: birthday,
		Gender:   input.Gender,
		PhotoURL: input.PhotoURL,
	}

	if errorCode := ctl.InfoQuery.Create(&userInfo); errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
