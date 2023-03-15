package user

import (
	"github.com/TskFok/GinApi/controller"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserName   string `validate:"required|min_len:5"`
	Password   string `validate:"required|min_len:5"`
	RePassword string `validate:"required|min_len:5"`
}

func registerValidate(ctx *gin.Context) bool {
	params := map[string]string{
		"user_name":   "required",
		"password":    "required|minLen:7",
		"re_password": "required|minLen:7",
	}

	return controller.CreateValidate(ctx, params)
}

func loginValidate(ctx *gin.Context) bool {
	params := map[string]string{
		"user_name": "required",
		"password":  "required|minLen:7",
	}

	return controller.CreateValidate(ctx, params)
}
