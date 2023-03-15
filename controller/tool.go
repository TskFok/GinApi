package controller

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetId(ctx *gin.Context) int {
	id := ctx.DefaultPostForm("id", "0")

	if id == "0" {
		id = ctx.DefaultQuery("id", "0")
	}

	intId, intErr := strconv.Atoi(id)

	if nil != intErr {
		response.Error(ctx, err.RuntimeError, intErr.Error())

		return 0
	}

	return intId
}
