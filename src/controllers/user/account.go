package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/ginsession"
	"yehuizhang.com/go-webapp-gin/src/dao/user/account"
)

func (ctl *Controller) SignIn(c *gin.Context) {

	input := account.SignInForm{}
	err := c.Bind(&input)

	if err != nil {
		ctl.Log.Errorw("unable to read account info from body", "err", err)
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

	addUidToSessionStore(c, record.Uuid.String())
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

	if err := addUidToSessionStore(c, id); err != nil {
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

func addUidToSessionStore(c *gin.Context, uid string) error {
	if store := ginsession.FromContext(c); store == nil {
		return fmt.Errorf("failed to get session store from context")
	} else {
		store.Set(UID, uid)
		store.Save()
	}
	return nil
}
