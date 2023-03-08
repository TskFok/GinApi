package router

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/tool"
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

	ctx.JSON(err.SUCCESS, tool.GetSuccess(list))

}

func Get(ctx *gin.Context) {
	id, exists := ctx.GetQuery("id")

	if !exists {
		ctx.JSON(err.UNDEFINED_ERROR, tool.GetErrorInfo(err.PARAMS_UNDEFINED_ERROR))
	}
	router := &model.Router{}

	condition := make(map[string]interface{})
	condition["id"] = id
	routerDetail, exists := router.Get(condition)

	if !exists {
		ctx.JSON(err.UNDEFINED_ERROR, tool.GetErrorInfo(err.ROUTER_UNDEFINED_ERROR))

		return
	}

	ctx.JSON(err.SUCCESS, tool.GetSuccess(routerDetail))
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
		ctx.JSON(err.UNDEFINED_ERROR, tool.GetErrorInfo(err.USER_UNDEFINED_ERROR))
	}
	newRouter.CreatorId = userId.(uint32)

	userName, exists := ctx.Get("user_name")

	if !exists {
		ctx.JSON(err.UNDEFINED_ERROR, tool.GetErrorInfo(err.USER_UNDEFINED_ERROR))
	}

	newRouter.CreatorName = userName.(string)

	id, routerErr := newRouter.Create(newRouter)

	if routerErr != nil {
		ctx.JSON(err.RUNTIME_ERROR, tool.GetErrorInfo(err.ROUTE_CREATE_ERROR))

		return
	}

	ctx.JSON(err.SUCCESS, tool.GetSuccess(id))
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
		ctx.JSON(err.UNDEFINED_ERROR, tool.GetErrorInfo(err.ROUTER_UNDEFINED_ERROR))

		return
	}

	condition = make(map[string]interface{})
	condition["router"] = router
	condition["description"] = description
	condition["method"] = method

	isUpdate := routerDetail.Update(condition)

	if !isUpdate {
		ctx.JSON(err.RUNTIME_ERROR, tool.GetErrorInfo(err.ROUTE_UPDATE_ERROR))

		return
	}
	ctx.JSON(err.SUCCESS, tool.GetSuccess("success"))
}
