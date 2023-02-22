package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
	router.Run(":8999")
}
