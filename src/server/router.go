package server

import (
	"github.com/gin-gonic/gin"
	"yehuizhang.com/go-webapp-gin/src/controllers"
	"yehuizhang.com/go-webapp-gin/src/database"
	"yehuizhang.com/go-webapp-gin/src/middlewares"
)

func NewRouter(database *database.Database) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Session(database))

	health := new(controllers.HealthController)
	user := new(controllers.UserController)

	router.GET("/health", health.Status)
	router.POST("/register", user.Signup)
	router.POST("/login", user.Signin)
	v1 := router.Group("v1")
	{
		v1.Use(middlewares.Auth)
		userGroup := v1.Group("user")
		{
			userGroup.GET("/info", user.GetInfo)
			userGroup.PUT("/info", user.UpdateInfo)
		}
	}

	return router
}
