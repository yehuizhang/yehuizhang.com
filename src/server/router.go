package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"yehuizhang.com/go-webapp-gin/pkg/database"
	"yehuizhang.com/go-webapp-gin/pkg/logger"
	"yehuizhang.com/go-webapp-gin/src/controllers"
	"yehuizhang.com/go-webapp-gin/src/controllers/admin"
	"yehuizhang.com/go-webapp-gin/src/controllers/user"
	"yehuizhang.com/go-webapp-gin/src/middlewares"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), controllers.ControllerSet)

type Router struct {
	AdminController *admin.Controller
	UserController  *user.Controller
	Redis           database.IRedis
	Log             *logger.Logger
}

func (r *Router) RegisterAPI(app *gin.Engine) {
	apiGroup := app.Group("/api")

	//apiGroup.Use(middlewares.ErrorHandler(r.Log))
	apiGroup.Use(gin.Logger())
	apiGroup.Use(gin.Recovery())
	apiGroup.Use(middlewares.Session(r.Redis))

	apiGroup.GET("/health", r.AdminController.GetHealth)
	v1 := apiGroup.Group("v1")
	{
		v1.POST("/register", r.UserController.SignUp)
		v1.POST("/login", r.UserController.SignIn)
		v1.Use(middlewares.Auth)
		userGroup := v1.Group("user")
		{
			userGroup.GET("/info", r.UserController.GetInfo)
			userGroup.POST("/info", r.UserController.CreateInfo)
			userGroup.PUT("/info", r.UserController.UpdateInfo)
		}
	}
}
