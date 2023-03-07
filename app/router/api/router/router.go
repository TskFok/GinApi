package router

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/tool"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	router := ctx.PostForm("router")
	description := ctx.PostForm("description")
	routerType := ctx.PostForm("type")

	newRouter := &model.Router{}

	newRouter.Router = router
	newRouter.Description = description
	newRouter.Type = routerType

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
