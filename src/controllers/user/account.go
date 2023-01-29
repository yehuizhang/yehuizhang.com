package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"yehuizhang.com/go-webapp-gin/src/dao/user/account"
	"yehuizhang.com/go-webapp-gin/src/utils/auth"
)

func (ctl *Controller) SignIn(c *gin.Context) {

	input, err := readCredentialFromContext(c)
	if err != nil {
		ctl.Log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	record, errorCode := ctl.AccountQuery.GetByUsername(input.Username)
	if errorCode != 0 {
		c.AbortWithStatus(errorCode)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(record.Password), []byte(input.Password))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	auth.AddUidToSessionStore(c, record.Uuid.String())
	c.String(http.StatusOK, "successfully logged in")
}

func (ctl *Controller) SignUp(c *gin.Context) {
	input, err := readCredentialFromContext(c)
	if err != nil {
		ctl.Log.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, errCode := ctl.AccountQuery.Create(input)
	if errCode != 0 {
		c.AbortWithStatus(errCode)
	}

	auth.AddUidToSessionStore(c, id)
	c.String(http.StatusCreated, "")
}

func readCredentialFromContext(c *gin.Context) (*account.Form, error) {
	form := account.Form{}
	err := c.Bind(&form)

	if err != nil {
		return nil, fmt.Errorf("unable to read account info from body: %s", err)
	}
	return &form, nil
}
