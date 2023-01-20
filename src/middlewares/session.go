package middlewares

import (
	"net/http"
	"yehuizhang.com/go-webapp-gin/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/go-session/redis/v3"
	"github.com/go-session/session/v3"
	"yehuizhang.com/go-webapp-gin/pkg/ginsession"
)

// Session Cookie expires in 7 days. Session is removed from DB in 2 hours, so it has to be refreshed to keep user live
func Session(database *database.Database) gin.HandlerFunc {
	sessionConfig := ginsession.Config{
		ErrorHandleFunc: func(c *gin.Context, err error) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		},
		StoreKey:  "",
		ManageKey: "",
		Skipper: func(c *gin.Context) bool {
			return false
		},
	}

	// This is the max length for user to keep the session without login again
	cookieLifeTime := 3600 * 24 * 2
	// If user idles longer than this time, session data will be removed from DB and user has to log in again. Keep active refreshes sessionLife
	sessionLifeTime := int64(3600 * 2)

	return ginsession.NewWithConfig(
		sessionConfig, session.SetStore(redis.NewRedisStoreWithCli(database.Redis, "user:session:")),
		session.SetSecure(false), session.SetSameSite(http.SameSiteLaxMode),
		session.SetCookieName("sid"), session.SetCookieLifeTime(cookieLifeTime),
		session.SetExpired(sessionLifeTime))
}
