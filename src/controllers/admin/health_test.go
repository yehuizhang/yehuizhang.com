package admin

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthController_Status(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	v := new(Controller)
	v.GetHealth(c)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
