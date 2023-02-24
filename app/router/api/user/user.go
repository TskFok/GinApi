package user

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	queryFilter := ctx.QueryMap("filter")
	filter := make(map[string]interface{})

	filter[queryFilter["name"]] = queryFilter["value"]

	test := &model.Test{}

	ctx.JSON(200, gin.H{
		"list": test.GetTest(filter),
	})
}

func Update(ctx *gin.Context) {
	id := ctx.PostForm("id")
	title := ctx.PostForm("title")

	model := &model.Test{}

	test := model.GetOne(id)

	filter := make(map[string]interface{})
	filter["title"] = title

	test.Update(filter)

	ctx.JSON(err.SUCCESS, gin.H{})
}
