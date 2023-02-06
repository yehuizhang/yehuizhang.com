package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Get health status of the server
// @Description	Get Health
// @Tags			Status
// @Produce		plain
// @Success		200	{string}	OK
// @Router			/health [get]
func (v *Controller) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
