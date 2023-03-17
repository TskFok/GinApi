package controller

import (
	"github.com/TskFok/GinApi/app/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetId(ctx *gin.Context) int {
	id := ctx.DefaultPostForm("id", "0")

	if id == "0" {
		id = ctx.DefaultQuery("id", "0")
	}

	intId, intErr := strconv.Atoi(id)

	if nil != intErr {
		response.Error(ctx, http.StatusBadRequest, intErr.Error())

		return 0
	}

	return intId
}
