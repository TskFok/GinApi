package router

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/response"
	"github.com/TskFok/GinApi/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func singleValidate(ctx *gin.Context) bool {
	params := map[string]string{
		"id": "required",
	}
	return controller.CreateValidate(ctx, params)
}

func createValidate(ctx *gin.Context) bool {
	params := map[string]string{
		"router":      "required",
		"method":      "required",
		"description": "max_len:50",
	}

	return controller.CreateValidate(ctx, params)
}

func updateValidate(ctx *gin.Context) bool {
	params := map[string]string{
		"id":          "required",
		"router":      "required",
		"method":      "required",
		"description": "max_len:50",
	}

	return controller.CreateValidate(ctx, params)
}

func exist(ctx *gin.Context, intId int) bool {
	routerModel := &model.Router{}
	routerModel.Router = ctx.PostForm("router")
	routerModel.Method = ctx.PostForm("method")

	router, exists := routerModel.Get(routerModel)

	if exists && router.Id != uint32(intId) {
		response.Error(ctx, http.StatusBadRequest, err.RouteRepeatError)

		return true
	}
	return false
}
