package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (v *Controller) GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
