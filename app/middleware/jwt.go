package middleware

import (
	"encoding/json"
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/response"
	"github.com/TskFok/GinApi/app/tool"
	"github.com/TskFok/GinApi/utils/cache"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		claims, tokenErr := tool.TokenInfo(token)

		if nil != tokenErr {
			response.Error(ctx, err.RUNTIME_ERROR, tokenErr.Error())

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
				response.Error(ctx, err.UNDEFINED_ERROR, err.USER_UNDEFINED_ERROR)

				ctx.Abort()
				return
			}
		} else {
			userModel := &model.User{}
			condition := make(map[string]interface{})
			condition["id"] = claims.Uid
			var exists bool

			user, exists = userModel.Get(condition)

			if !exists {
				response.Error(ctx, err.UNDEFINED_ERROR, err.USER_UNDEFINED_ERROR)

				ctx.Abort()
				return
			}
			res, jsonErr := json.Marshal(user)

			if nil != jsonErr {
				response.Error(ctx, err.RUNTIME_ERROR, err.REDIS_ERROR)

				ctx.Abort()
				return
			}
			cache.Set(key, string(res), 3600)
		}

		ctx.Set("user", user)
		ctx.Set("user_name", user.UserName)
		ctx.Set("user_id", user.Id)

		ctx.Next()
	}
}
