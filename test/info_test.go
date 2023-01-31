package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
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

	c.Request = httptest.NewRequest(http.MethodPost, "/", toReader(input))
	c.Set(user.UID, "test")
	v := user.Controller{InfoQuery: &mockedInfoQuery}
	v.Create(c)

	assert.Equal(t, 200, w.Code)

}

func TestController_Get(t *testing.T) {
	// TODO
}
