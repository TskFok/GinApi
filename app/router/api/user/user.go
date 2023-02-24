package user

import (
	"github.com/TskFok/GinApi/app/model"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	id := ctx.Query("id")
	filter := make(map[string]interface{})

	filter["id"] = id

	ctx.JSON(200, gin.H{
		"list": model.GetTest(filter),
	})
}
