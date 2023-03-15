package controller

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

func CreateValidate(ctx *gin.Context, params map[string]string) bool {
	face, validateError := validate.FromRequest(ctx.Request)

	if nil != validateError {
		response.Error(ctx, err.RuntimeError, err.PasswordValidateError)

		return false
	}

	v := face.Create()

	for field, rule := range params {
		v.StringRule(field, rule)
	}

	if !v.Validate() {
		response.Error(ctx, err.RuntimeError, v.Errors.One())

		return false
	}

	return true
}
