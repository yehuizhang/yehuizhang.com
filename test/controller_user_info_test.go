package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
	"yehuizhang.com/go-webapp-gin/pkg/dao/user/info"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
)

func TestInputForm_Invalid_Birthday(t *testing.T) {
	c, _ := createGinContext()

	data := map[string]interface{}{
		"name":     "name",
		"birthday": "1988",
		"gender":   "F",
		"photoURL": "",
	}

	inputX := &info.Form{}

	c.Request = generateRequest(http.MethodPost, "/", data)
	err := c.ShouldBind(inputX)
	assert.NotNil(t, err)
}

func TestController_CreateInfo(t *testing.T) {
	c, w := createGinContext()

	input := map[string]interface{}{
		"name":     "name",
		"birthday": "1988-01-02T08:00:00.000Z",
		"gender":   "F",
		"photoURL": "",
	}
	mockedInfoQuery := IUserInfoQuery{}
	mockedInfoQuery.On("Create", mock.Anything).Return(0)

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")
	v := user.Controller{InfoQuery: &mockedInfoQuery, Log: lg}
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

	input := map[string]interface{}{
		"name":     "name",
		"birthday": "1988",
		"gender":   "F",
		"photoURL": "",
	}

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")

	mockedInfoQuery := IUserInfoQuery{}
	v := user.Controller{InfoQuery: &mockedInfoQuery, Log: lg}
	v.CreateInfo(c)

	assert.Equal(t, 400, c.Writer.Status())
}

func TestController_CreateInfo_DB_Failed(t *testing.T) {
	c, w := createGinContext()

	data := map[string]interface{}{
		"name":     "name",
		"birthday": "1988-01-01",
		"gender":   "F",
		"photoURL": "",
	}
	jsonData, _ := json.Marshal(data)
	mockedInfoQuery := IUserInfoQuery{}
	mockedInfoQuery.On("Create", mock.Anything).Return(400)

	c.Request = generateRequest(http.MethodPost, "/", bytes.NewReader(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")
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
		Id:        "uuid-test",
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
	assert.Equal(t, mockedResult.Id, output.Id)
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
