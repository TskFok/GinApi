package router

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/response"
	"github.com/TskFok/GinApi/controller"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if !singleValidate(ctx) {
		return
	}

	id, exists := ctx.GetQuery("id")

	if !exists {
		response.Error(ctx, http.StatusNotFound, err.ParamsUndefinedError)

		return
	}
	router := &model.Router{}

	condition := make(map[string]interface{})
	condition["id"] = id
	routerDetail, exists := router.Get(condition)

	if !exists {
		response.Error(ctx, http.StatusNotFound, err.RouterUndefinedError)

		return
	}

	response.Success(ctx, routerDetail)
}

func Create(ctx *gin.Context) {
	if !createValidate(ctx) {
		return
	}

	newRouter := &model.Router{}

	newRouter.Router = ctx.PostForm("router")
	newRouter.Method = ctx.PostForm("method")

	_, exists := newRouter.Get(newRouter)

	if exists {
		response.Error(ctx, http.StatusNotFound, err.RouteHasExistsError)

		return
	}

	newRouter.Description = ctx.PostForm("description")

	userId, exists := ctx.Get("user_id")

	if !exists {
		response.Error(ctx, http.StatusNotFound, err.UserUndefinedError)

		return
	}
	newRouter.CreatorId = userId.(uint32)

	userName, exists := ctx.Get("user_name")

	if !exists {
		response.Error(ctx, http.StatusNotFound, err.UserUndefinedError)

		return
	}

	newRouter.CreatorName = userName.(string)

	id, routerErr := newRouter.Create(newRouter)

	if routerErr != nil {
		response.Error(ctx, http.StatusBadRequest, err.RouteCreateError)

		return
	}

	response.Success(ctx, id)
}

func Update(ctx *gin.Context) {
	if !updateValidate(ctx) {
		return
	}

	intId := controller.GetId(ctx)

	if exist(ctx, intId) {
		return
	}

	routerModel := &model.Router{}
	routerModel.Id = uint32(intId)

	router, exists := routerModel.Get(routerModel)

	if !exists {
		response.Error(ctx, http.StatusNotFound, err.RouteNotExistsError)

		return
	}

	router.Description = ctx.PostForm("description")
	router.Router = ctx.PostForm("router")
	router.Method = ctx.PostForm("method")

	isUpdate := routerModel.Update(router)

	if !isUpdate {
		response.Error(ctx, http.StatusBadRequest, err.RouteUpdateError)

		return
	}

	response.Success(ctx, "success")
}
