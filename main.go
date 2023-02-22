package main

import (
	"fmt"
	"github.com/TskFok/GinApi/app/utils/conf"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	readTimeOut := conf.GetConf("app.read_time_out")
	writeTimeOut := conf.GetConf("app.write_time_out")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8999),
		Handler:        router,
		ReadTimeout:    time.Duration(readTimeOut.(int)) * time.Second,
		WriteTimeout:   time.Duration(writeTimeOut.(int)) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
