package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"testing"
	"yehuizhang.com/go-webapp-gin/pkg/dao/shared"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/account"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
)

func TestController_SignUp(t *testing.T) {
	c, w := createGinContext()

	input := account.SignUpForm{
		Username: "username",
		Password: "password",
		Email:    "test@test.com",
	}
	mockedAccountQuery := IUserAccountQuery{}
	mockedAccountQuery.On("Create", mock.Anything).Return("test-id", 0)
	c.Request = generateRequest(http.MethodPost, "/", input)
	v := user.Controller{AccountQuery: &mockedAccountQuery, Log: lg}
	v.SignUp(c)
	assert.Equal(t, 201, w.Code)
}

func TestController_SignUp_Invalid_Input(t *testing.T) {
	c, w := createGinContext()
	mockedAccountQuery := IUserAccountQuery{}

	input := "invalid value"
	c.Request = generateRequest(http.MethodPost, "/", input)
	v := user.Controller{AccountQuery: &mockedAccountQuery, Log: lg}
	v.SignUp(c)
	assert.Equal(t, 400, w.Code)
}

func TestController_SignUp_Db_failure(t *testing.T) {
	c, w := createGinContext()

	input := account.SignUpForm{
		Username: "username",
		Password: "password",
		Email:    "test@test.com",
	}
	mockedAccountQuery := IUserAccountQuery{}
	mockedAccountQuery.On("Create", mock.Anything).Return("", 777)
	c.Request = generateRequest(http.MethodPost, "/", input)
	v := user.Controller{AccountQuery: &mockedAccountQuery, Log: lg}
	v.SignUp(c)
	assert.Equal(t, 777, w.Code)
}

func TestController_SignIn(t *testing.T) {
	c, w := createGinContext()

	input := account.SignInForm{
		Username: "username",
		Password: "password",
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 4)
	dbResult := account.UserAccount{
		Model:    shared.Model{},
		Username: "username",
		Password: string(encryptedPassword),
		Email:    "",
		Active:   false,
	}
	mockedAccountQuery := IUserAccountQuery{}
	mockedAccountQuery.On("GetByUsername", mock.Anything).Return(&dbResult, 0)
	c.Request = generateRequest(http.MethodPost, "/", input)
	v := user.Controller{AccountQuery: &mockedAccountQuery, Log: lg}
	v.SignIn(c)
	assert.Equal(t, 200, w.Code)
}

func TestController_SignIn_Invalid_Password(t *testing.T) {
	c, w := createGinContext()

	input := account.SignInForm{
		Username: "username",
		Password: "password",
	}

	dbResult := account.UserAccount{
		Model:    shared.Model{},
		Username: "username",
		Password: "password",
		Email:    "",
		Active:   false,
	}
	mockedAccountQuery := IUserAccountQuery{}
	mockedAccountQuery.On("GetByUsername", mock.Anything).Return(&dbResult, 0)
	c.Request = generateRequest(http.MethodPost, "/", input)
	v := user.Controller{AccountQuery: &mockedAccountQuery, Log: lg}
	v.SignIn(c)
	assert.Equal(t, 401, w.Code)
}

func TestController_SignIn_DB_Failure(t *testing.T) {
	c, w := createGinContext()

	input := account.SignInForm{
		Username: "username",
		Password: "password",
	}
	mockedAccountQuery := IUserAccountQuery{}
	mockedAccountQuery.On("GetByUsername", mock.Anything).Return(nil, 777)
	c.Request = generateRequest(http.MethodPost, "/", input)
	v := user.Controller{AccountQuery: &mockedAccountQuery, Log: lg}
	v.SignIn(c)
	assert.Equal(t, 777, w.Code)
}

func TestController_SignIn_Invalid_Input(t *testing.T) {
	c, w := createGinContext()
	mockedAccountQuery := IUserAccountQuery{}

	input := "invalid"
	c.Request = generateRequest(http.MethodPost, "/", input)
	v := user.Controller{AccountQuery: &mockedAccountQuery, Log: lg}
	v.SignIn(c)
	assert.Equal(t, 400, w.Code)
}

//
//func TestController_CreateInfo_DB_Failed(t *testing.T) {
//	c, w := createGinContext()
//
//	input := info.Form{
//		Name:     "name",
//		Birthday: "1988-01-01",
//		Gender:   "F",
//		PhotoURL: "",
//	}
//	mockedInfoQuery := IUserInfoQuery{}
//	mockedInfoQuery.On("Create", mock.Anything).Return(400)
//
//	c.Request = generateRequest(http.MethodPost, "/", input)
//	c.Set(user.UID, "test")
//	v := user.Controller{InfoQuery: &mockedInfoQuery}
//	v.CreateInfo(c)
//	assert.Equal(t, 400, w.Code)
//}
//
//func TestController_GetInfo(t *testing.T) {
//	c, w := createGinContext()
//
//	c.Request = generateRequest(http.MethodGet, "/", "")
//	c.Set(user.UID, "test")
//
//	mockedResult := info.UserInfo{
//		Uuid:      "uuid-test",
//		Name:      "",
//		Birthday:  time.Time{},
//		Gender:    "",
//		PhotoURL:  "",
//		CreatedAt: time.Time{},
//		UpdatedAt: time.Time{},
//	}
//
//	mockedInfoQuery := IUserInfoQuery{}
//	mockedInfoQuery.On("Get", mock.Anything).Return(&mockedResult, 0)
//
//	v := user.Controller{InfoQuery: &mockedInfoQuery}
//	v.GetInfo(c)
//	assert.Equal(t, 200, w.Code)
//
//	output := info.UserInfo{}
//	err := json.Unmarshal(w.Body.Bytes(), &output)
//	if err != nil {
//		t.FailNow()
//	}
//	assert.Equal(t, mockedResult.Uuid, output.Uuid)
//}
//
//func TestController_GetInfo_DB_Failed(t *testing.T) {
//	c, w := createGinContext()
//
//	c.Request = generateRequest(http.MethodGet, "/", "")
//	c.Set(user.UID, "test")
//
//	mockedInfoQuery := IUserInfoQuery{}
//	mockedInfoQuery.On("Get", mock.Anything).Return(nil, 400)
//
//	v := user.Controller{InfoQuery: &mockedInfoQuery}
//	v.GetInfo(c)
//	assert.Equal(t, 400, w.Code)
//}
