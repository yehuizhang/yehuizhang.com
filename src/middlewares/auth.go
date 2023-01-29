package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yehuizhang.com/go-webapp-gin/pkg/ginsession"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
)

// validate user's identify and set into context
func Auth(c *gin.Context) {

	store := ginsession.FromContext(c)

	uid, ok := store.Get(user.UID)

	if ok {
		c.Set(user.UID, uid)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
