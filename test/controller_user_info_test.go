package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
	"yehuizhang.com/go-webapp-gin/src/dao/user/info"
)

func TestController_Create(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

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
	v.Create(c)
	assert.Equal(t, 200, w.Code)

	output := info.UserInfo{}
	err := json.Unmarshal(w.Body.Bytes(), &output)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, "name", output.Name)
}

func TestController_Create_Invalid_Input(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	input := "invalid value"

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")

	mockedInfoQuery := IUserInfoQuery{}
	v := user.Controller{InfoQuery: &mockedInfoQuery, Log: logger.InitLogger(flagParser)}
	v.Create(c)

	assert.Equal(t, 500, c.Writer.Status())
}

func TestController_Create_Invalid_Birthday(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	input := info.Form{
		Name:     "name",
		Birthday: "1988",
		Gender:   "F",
		PhotoURL: "",
	}

	c.Request = generateRequest(http.MethodPost, "/", input)
	c.Set(user.UID, "test")

	mockedInfoQuery := IUserInfoQuery{}
	v := user.Controller{InfoQuery: &mockedInfoQuery, Log: logger.InitLogger(flagParser)}
	v.Create(c)

	assert.Equal(t, 500, c.Writer.Status())
}

func TestController_Create_DB_Failed(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

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
	v.Create(c)
	assert.Equal(t, 400, w.Code)
}

func TestController_Get(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

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
	v.Get(c)
	assert.Equal(t, 200, w.Code)

	output := info.UserInfo{}
	err := json.Unmarshal(w.Body.Bytes(), &output)
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, mockedResult.Uuid, output.Uuid)
}

func TestController_Get_DB_Failed(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = generateRequest(http.MethodGet, "/", "")
	c.Set(user.UID, "test")

	mockedInfoQuery := IUserInfoQuery{}
	mockedInfoQuery.On("Get", mock.Anything).Return(nil, 400)

	v := user.Controller{InfoQuery: &mockedInfoQuery}
	v.Get(c)
	assert.Equal(t, 400, w.Code)
}
