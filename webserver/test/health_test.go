package test

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
	"yehuizhang.com/go-webapp-gin/src/controllers/admin"

	"github.com/stretchr/testify/assert"
)

func TestHealthController_Status(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	v := new(admin.Controller)
	v.GetHealth(c)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
