package middlewares

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
	"yehuizhang.com/go-webapp-gin/pkg/ginsession"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
)

// Auth validate user's identify and set into context
func Auth(c *gin.Context) {

	store := ginsession.FromContext(c)

	if uid, ok := store.Get(user.UID); ok {
		if uid, ok := uid.(string); ok {
			if _, err := uuid.Parse(uid); err == nil {
				c.Set(user.UID, uid)
				return
			}
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
