package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/src/controllers"
	"yehuizhang.com/go-webapp-gin/src/middlewares"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), controllers.ControllerSet)

type Router struct {
	HealthController   *controllers.HealthController
	UserAuthController *controllers.UserAuthController
	UserInfoController *controllers.UserInfoController
	Database           *database.Database
}

func (r *Router) RegisterAPI(app *gin.Engine) {
	apiGroup := app.Group("/api")

	apiGroup.Use(gin.Logger())
	apiGroup.Use(gin.Recovery())
	apiGroup.Use(middlewares.Session(r.Database))

	apiGroup.GET("/health", r.HealthController.Get)
	v1 := apiGroup.Group("v1")
	{

		v1.POST("/register", r.UserAuthController.SignUp)
		v1.POST("/login", r.UserAuthController.SignIn)

		v1.Use(middlewares.Auth)
		userGroup := v1.Group("user")
		{
			userGroup.GET("/info", r.UserInfoController.Get)
			userGroup.PUT("/info", r.UserInfoController.CreateOrUpdate)
		}
	}
}
