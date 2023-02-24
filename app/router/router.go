package router

import (
	"fmt"
	"github.com/TskFok/GinApi/app/middleware"
	"github.com/TskFok/GinApi/app/router/api/user"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	gin.SetMode(conf.AppRunMode)

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.HTML(200, message, gin.H{
			"status": "success",
		})
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"all": "hi",
		})
	})

	api := router.Group("/api")
	{
		userApi := api.Group("/user")
		{
			userApi.POST("/login", user.Login)
			userApi.POST("/register", user.Register)
			userApi.Use(middleware.Jwt())
			userApi.GET("/info", user.Info)
		}
	}

	return router
}
