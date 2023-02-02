package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
	"yehuizhang.com/go-webapp-gin/src/dao/user/info"
)

func TestController_CreateInfo(t *testing.T) {
	c, w := createGinContext()

	input := info.Form{
		Name:     "name",
		Birthday: "1988-01-01",
		Gender:   "F",
		PhotoURL: "",
	}
	mockedInfoQuery := IUserInfoQuery{}
	mockedInfoQuery.On("Create", mock.Anything).Return(0)

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")
	v := user.Controller{InfoQuery: &mockedInfoQuery}
	v.CreateInfo(c)
	assert.Equal(t, 200, w.Code)

	output := info.UserInfo{}
	err := json.Unmarshal(w.Body.Bytes(), &output)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, "name", output.Name)
}

func TestController_CreateInfo_Invalid_Input(t *testing.T) {
	c, _ := createGinContext()

	input := "invalid value"

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")

	mockedInfoQuery := IUserInfoQuery{}
	v := user.Controller{InfoQuery: &mockedInfoQuery, Log: lg}
	v.CreateInfo(c)

	assert.Equal(t, 500, c.Writer.Status())
}

func TestController_CreateInfo_Invalid_Birthday(t *testing.T) {
	c, _ := createGinContext()

	input := info.Form{
		Name:     "name",
		Birthday: "1988",
		Gender:   "F",
		PhotoURL: "",
	}

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")

	mockedInfoQuery := IUserInfoQuery{}
	v := user.Controller{InfoQuery: &mockedInfoQuery, Log: lg}
	v.CreateInfo(c)

	assert.Equal(t, 500, c.Writer.Status())
}

func TestController_CreateInfo_DB_Failed(t *testing.T) {
	c, w := createGinContext()

	input := info.Form{
		Name:     "name",
		Birthday: "1988-01-01",
		Gender:   "F",
		PhotoURL: "",
	}
	mockedInfoQuery := IUserInfoQuery{}
	mockedInfoQuery.On("Create", mock.Anything).Return(400)

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")
	v := user.Controller{InfoQuery: &mockedInfoQuery}
	v.CreateInfo(c)
	assert.Equal(t, 400, w.Code)
}

func TestController_GetInfo(t *testing.T) {
	c, w := createGinContext()

	c.Request = generateRequest(http.MethodGet, "/", "")
	c.Set(user.UID, "test")

	mockedResult := info.UserInfo{
		Uuid:      "uuid-test",
		Name:      "",
		Birthday:  time.Time{},
		Gender:    "",
		PhotoURL:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mockedInfoQuery := IUserInfoQuery{}
	mockedInfoQuery.On("Get", mock.Anything).Return(&mockedResult, 0)

	v := user.Controller{InfoQuery: &mockedInfoQuery}
	v.GetInfo(c)
	assert.Equal(t, 200, w.Code)

	output := info.UserInfo{}
	err := json.Unmarshal(w.Body.Bytes(), &output)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, mockedResult.Uuid, output.Uuid)
}

func TestController_GetInfo_DB_Failed(t *testing.T) {
	c, w := createGinContext()

	c.Request = generateRequest(http.MethodGet, "/", "")
	c.Set(user.UID, "test")

	mockedInfoQuery := IUserInfoQuery{}
	mockedInfoQuery.On("Get", mock.Anything).Return(nil, 400)

	v := user.Controller{InfoQuery: &mockedInfoQuery}
	v.GetInfo(c)
	assert.Equal(t, 400, w.Code)
}
