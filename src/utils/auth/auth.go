package auth

import (
	"log"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"

	"github.com/gin-gonic/gin"
	"yehuizhang.com/go-webapp-gin/pkg/ginsession"
)

func AddUidToSessionStore(c *gin.Context, uid string) {
	store := ginsession.FromContext(c)

	log.Printf("store: %t", store == nil)

	store.Set(user.UID, uid)
	store.Save()
}
