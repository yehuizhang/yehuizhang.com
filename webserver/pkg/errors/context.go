package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddNewContextError(c *gin.Context, err error) error {
	c.AbortWithStatus(http.StatusBadRequest)
	return fmt.Errorf("failed to add value to context. %s", err)
}
