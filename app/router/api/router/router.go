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

	id, routerErr := newRouter.Create(newRouter)

	if routerErr != nil {
		ctx.JSON(err.RUNTIME_ERROR, tool.GetErrorInfo(err.ROUTE_CREATE_ERROR))

		return
	}

	ctx.JSON(err.SUCCESS, tool.GetSuccess(id))
}
