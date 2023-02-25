package middlewares

import (
	"github.com/gin-gonic/gin"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
)

func ErrorHandler(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			// TODO: This is not usable until error stack is fixed
			logger.Errorf("%+v", err.Error())
		}
	}
}
