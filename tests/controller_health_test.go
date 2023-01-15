package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"yehuizhang.com/go-webapp-gin/controllers"
)

func TestHealthController_Status(t *testing.T) {

	router := SetupRouter()
	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
