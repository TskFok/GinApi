package router

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func List(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	size := ctx.DefaultQuery("size", "10")

	router := &model.Router{}

	pageInt, _ := strconv.Atoi(page)
	sizeInt, _ := strconv.Atoi(size)

	list := router.List(pageInt, sizeInt)

	response.Success(ctx, list)

}

func Get(ctx *gin.Context) {
	id, exists := ctx.GetQuery("id")

	if !exists {
		response.Error(ctx, err.UNDEFINED_ERROR, err.PARAMS_UNDEFINED_ERROR)

		return
	}
	router := &model.Router{}

	condition := make(map[string]interface{})
	condition["id"] = id
	routerDetail, exists := router.Get(condition)

	if !exists {
		response.Error(ctx, err.UNDEFINED_ERROR, err.ROUTER_UNDEFINED_ERROR)

		return
	}

	response.Success(ctx, routerDetail)
}

func Create(ctx *gin.Context) {
	router := ctx.PostForm("router")
	description := ctx.PostForm("description")
	method := ctx.PostForm("method")

	newRouter := &model.Router{}

	newRouter.Router = router
	newRouter.Description = description
	newRouter.Method = method

	userId, exists := ctx.Get("user_id")

	if !exists {
		response.Error(ctx, err.UNDEFINED_ERROR, err.USER_UNDEFINED_ERROR)

		return
	}
	newRouter.CreatorId = userId.(uint32)

	userName, exists := ctx.Get("user_name")

	if !exists {
		response.Error(ctx, err.UNDEFINED_ERROR, err.USER_UNDEFINED_ERROR)

		return
	}

	newRouter.CreatorName = userName.(string)

	id, routerErr := newRouter.Create(newRouter)

	if routerErr != nil {
		response.Error(ctx, err.RUNTIME_ERROR, err.ROUTE_CREATE_ERROR)

		return
	}

	response.Success(ctx, id)
}

func Update(ctx *gin.Context) {
	id := ctx.PostForm("id")
	router := ctx.PostForm("router")
	description := ctx.PostForm("description")
	method := ctx.PostForm("method")

	routerModel := &model.Router{}
	condition := make(map[string]interface{})
	condition["id"] = id

	routerDetail, exists := routerModel.Get(condition)

	if !exists {
		response.Error(ctx, err.UNDEFINED_ERROR, err.ROUTER_UNDEFINED_ERROR)

		return
	}

	condition = make(map[string]interface{})
	condition["router"] = router
	condition["description"] = description
	condition["method"] = method

	isUpdate := routerDetail.Update(condition)

	if !isUpdate {
		response.Error(ctx, err.RUNTIME_ERROR, err.ROUTE_UPDATE_ERROR)

		return
	}

	response.Success(ctx, "success")
}
