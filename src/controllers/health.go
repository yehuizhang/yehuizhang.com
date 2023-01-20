package controllers

import (
	"github.com/google/wire"
	"net/http"

	"github.com/gin-gonic/gin"
)

var HealthControllerSet = wire.NewSet(wire.Struct(new(HealthController), "*"))

type HealthController struct{}

func (h *HealthController) Get(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
