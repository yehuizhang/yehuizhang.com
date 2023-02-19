package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/account"
	"yehuizhang.com/go-webapp-gin/pkg/ginsession"
)

func (ctl *Controller) SignIn(c *gin.Context) {

	//ctx, err := ctxUtil.NewTransactionLockShare(c.Request.Context())
	//if err != nil {
	//	err := errors.AddNewContextError(c, err)
	//	ctl.Log.Error(err)
	//	return
	//}

	input := account.SignInForm{}
	err := c.Bind(&input)

	if err != nil {
		ctl.Log.Errorw("unable to read account info from body", "err", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	record, errorCode := ctl.AccountQuery.GetByUsername(c.Request.Context(), input.Username)
	if errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(record.Password), []byte(input.Password))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := ginsession.AddValueToStore(c, UID, record.Id.String()); err != nil {
		ctl.Log.Error(err)
	}
	c.String(http.StatusOK, "successfully logged in")
}

// @formatter:off
// @Summary		Register new account
// @Description	Register new account
// @Tags			User
// @Accept			json
// @Param			request	body	account.SignUpForm	true	"account credential"
// @Produce		plain
// @Success		201
// @Router			/register [post]
func (ctl *Controller) SignUp(c *gin.Context) {
	input, err := readCredentialFromContext(c)
	if err != nil {
		ctl.Log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, errCode := ctl.AccountQuery.Create(c.Request.Context(), input)
	if errCode != 0 {
		c.AbortWithStatus(errCode)
		return
	}

	if err := ginsession.AddValueToStore(c, UID, id); err != nil {
		ctl.Log.Error(err)
	}
	c.String(http.StatusCreated, "")
}

func readCredentialFromContext(c *gin.Context) (*account.SignUpForm, error) {
	form := account.SignUpForm{}
	err := c.Bind(&form)

	if err != nil {
		return nil, fmt.Errorf("unable to read account info from body: %s", err)
	}
	return &form, nil
}
