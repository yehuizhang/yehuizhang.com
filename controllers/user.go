package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"yehuizhang.com/go-webapp-gin/forms"
	"yehuizhang.com/go-webapp-gin/models/user"
	"yehuizhang.com/go-webapp-gin/util/auth"
)

type UserController struct{}

var userCredentialModel = new(user.UserCredential)
var userInfoModel = new(user.UserInfo)

func (u UserController) Signin(c *gin.Context) {

	userCredential, err := userCredentialModel.Signin(c)

	if err != nil {
		if err == user.PasswordNotMatch {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Password does not match.",
			})
		} else {
			c.Status(http.StatusInternalServerError)
		}
		c.Abort()
		return
	}

	userInfo, err := user.GetUserInfo(userCredential.ID)
	if err != nil {
		log.Print("Unable to get UserInfo.", err)
		c.Abort()
		return
	}

	auth.AddUidToSessionStore(c, userCredential.ID)

	c.JSON(http.StatusOK, userInfo)
}

func (u UserController) Signup(c *gin.Context) {
	userInfo, userCredential, err := userCredentialModel.Signup(c)

	if err != nil {
		log.Print("signup failed: ", err)
		c.Abort()
		return
	}

	auth.AddUidToSessionStore(c, userCredential.ID)

	log.Println("Successfully signed up for user", *userCredential)
	c.JSON(http.StatusCreated, *userInfo)
}

func (u UserController) UpdateInfo(c *gin.Context) {
	uid := c.GetString(user.UID)
	userInfoInput := forms.UserInfo{}
	err := c.Bind(&userInfoInput)

	if err != nil {
		log.Fatal("Unable to read UserInfo from body: ", err)
		c.Abort()
		return
	}

	var userInfo *user.UserInfo
	userInfo, err = userInfoModel.AddOrUpdate(uid, userInfoInput)

	if err != nil {
		log.Print("Unable to update UserInfo", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, *userInfo)
}

func (u UserController) GetInfo(c *gin.Context) {

	uid, ok := c.Get(user.UID)

	if !ok {
		log.Print("uid was not found", user.UID)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userInfo, err := user.GetUserInfo(uid.(string))

	if err != nil {
		log.Print("Unable to get UserInfo.", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, userInfo)

}
