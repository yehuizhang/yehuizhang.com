package test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"yehuizhang.com/go-webapp-gin/pkg/flags"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

var fg = &flags.FlagParser{
	Env:        "test",
	ConfigPath: "",
	ConfigName: "",
	ConfigType: "",
}

var lg = logger.InitLogger(fg)

func toReader(v interface{}) io.Reader {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil
	}
	return &buf
}

func generateRequest(method string, target string, jsonBody interface{}) *http.Request {
	request := httptest.NewRequest(method, target, toReader(jsonBody))
	request.Header.Set("Content-Type", "application/json")
	return request
}

func createGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	return c, w
}
