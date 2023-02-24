package middleware

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/tool"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		claims, error := tool.TokenInfo(token)

		if nil != error {
			ctx.JSON(err.RUNTIME_ERROR, gin.H{
				"code": err.RUNTIME_ERROR,
				"msg":  error.Error(),
				"data": make(map[string]interface{}),
			})

			ctx.Abort()
			return
		}

		userModel := &model.User{}
		condition := make(map[string]interface{})
		condition["id"] = claims.Uid

		user, exists := userModel.HasOneByName(condition)

		if !exists {
			ctx.JSON(err.UNDEFINED_ERROR, err.GetErrorInfo(err.USER_UNDEFINED_ERROR))
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
