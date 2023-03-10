package router

import (
	"github.com/TskFok/GinApi/app/middleware"
	"github.com/TskFok/GinApi/app/router/api/router"
	"github.com/TskFok/GinApi/app/router/api/user"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/gin-gonic/gin"
)

var Handle *gin.Engine

func init() {
	Handle = gin.New()
	Handle.Use(gin.Recovery())
	Handle.Use(gin.Logger())

	gin.SetMode(conf.AppRunMode)

	api := Handle.Group("/api")
	{
		userApi := api.Group("/user")
		{
			userApi.POST("/login", user.Login)
			userApi.POST("/register", user.Register)
			userApi.Use(middleware.Jwt())
			userApi.GET("/info", user.Info)
		}
		routerApi := api.Group("/router")

		routerApi.Use(middleware.Jwt())
		{
			routerApi.GET("", router.List)
			routerApi.GET("/detail", router.Get)
			routerApi.POST("", router.Create)
			routerApi.PUT("", router.Update)
		}
	}
}
