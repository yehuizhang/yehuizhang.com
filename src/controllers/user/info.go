package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
	"yehuizhang.com/go-webapp-gin/pkg/errors"
	"yehuizhang.com/go-webapp-gin/src/util/ctxUtil"
)

// @formatter:off
// @Summary		Get User Info
// @Description	Get User Info
// @Tags			User
// @Produce		json
// @Success		200	{object}	info.UserInfo
// @Router			/user/info [get]
func (ctl *Controller) GetInfo(c *gin.Context) {
	uid := c.GetString(UID)

	userInfo, errorCode := ctl.InfoQuery.Get(c.Request.Context(), uid)

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

	if errorCode := ctl.InfoQuery.Create(c.Request.Context(), &userInfo); errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

func (ctl *Controller) UpdateInfo(c *gin.Context) {
	ctx, err := ctxUtil.NewTransactionLockShare(c.Request.Context())
	if err != nil {
		err := errors.AddNewContextError(c, err)
		ctl.Log.Error(err)
		return
	}

	uid := c.GetString(UID)
	input := info.Form{}
	if err := c.ShouldBind(&input); err != nil {
		ctl.Log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctl.Transaction.Exec(ctx, func(ctx context.Context) error {
		userInfo, errorCode := ctl.InfoQuery.Get(ctx, uid)
		if errorCode != 0 {
			c.AbortWithStatus(errorCode)
			return fmt.Errorf("unable to get userInfo from db for user %s", uid)
		}

		userInfo.Name = input.Name
		userInfo.Birthday = input.Birthday
		userInfo.Gender = input.Gender
		userInfo.PhotoURL = input.PhotoURL

		if errorCode := ctl.InfoQuery.Update(ctx, userInfo); errorCode != 0 {
			c.AbortWithStatus(errorCode)
			return fmt.Errorf("failed to update userInfo")
		}
		c.JSON(http.StatusOK, userInfo)
		return nil
	})

}
