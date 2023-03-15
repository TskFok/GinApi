package router

import (
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/app/middleware"
	"github.com/TskFok/GinApi/controller/router"
	"github.com/TskFok/GinApi/controller/user"
	"github.com/gin-gonic/gin"
)

var Handle *gin.Engine

func InitRouter() {
	gin.SetMode(global.Env)

	Handle = gin.New()
	Handle.Use(gin.Recovery())
	Handle.Use(gin.Logger())

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
