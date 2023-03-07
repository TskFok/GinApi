package middleware

import (
	"encoding/json"
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/tool"
	"github.com/TskFok/GinApi/app/utils/cache"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		claims, tokenErr := tool.TokenInfo(token)

		if nil != tokenErr {
			ctx.JSON(err.RUNTIME_ERROR, gin.H{
				"code": err.RUNTIME_ERROR,
				"msg":  tokenErr.Error(),
				"data": make(map[string]interface{}),
			})

			ctx.Abort()
			return
		}

		builder := strings.Builder{}
		builder.WriteString("user:info:")
		builder.WriteString(strconv.FormatUint(uint64(claims.Uid), 10))
		key := builder.String()

		user := &model.User{}
		if cache.Has(key) {
			jsonErr := json.Unmarshal([]byte(cache.Get(key)), &user)

			if nil != jsonErr {
				ctx.JSON(err.UNDEFINED_ERROR, tool.GetErrorInfo(err.USER_UNDEFINED_ERROR))
			}
		} else {
			userModel := &model.User{}
			condition := make(map[string]interface{})
			condition["id"] = claims.Uid
			var exists bool

			user, exists = userModel.HasOneByName(condition)

			if !exists {
				ctx.JSON(err.UNDEFINED_ERROR, tool.GetErrorInfo(err.USER_UNDEFINED_ERROR))
			}
			res, jsonErr := json.Marshal(user)

			if nil != jsonErr {
				ctx.JSON(err.RUNTIME_ERROR, tool.GetErrorInfo(err.REDIS_ERROR))
			}
			cache.Set(key, string(res), 3600)
		}

		ctx.Set("user", user)
		ctx.Set("user_name", user.UserName)
		ctx.Set("user_id", user.Id)

		ctx.Next()
	}
}
