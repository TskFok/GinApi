package user

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/response"
	"github.com/TskFok/GinApi/app/tool"
	"github.com/TskFok/GinApi/utils/cache"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func Info(ctx *gin.Context) {
	if user, exists := ctx.Get("user"); exists {
		response.Success(ctx, user)

		return
	}

	response.Error(ctx, err.UNDEFINED_ERROR, err.USER_UNDEFINED_ERROR)
}

func Login(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")

	userModel := &model.User{}

	condition := make(map[string]interface{})
	condition["user_name"] = userName

	user, exists := userModel.Get(condition)

	if !exists {
		response.Error(ctx, err.UNDEFINED_ERROR, err.USER_UNDEFINED_ERROR)

		return
	}

	if tool.Password(password, user.Salt) == user.Password {
		data := make(map[string]interface{})

		condition = make(map[string]interface{})
		condition["last_login_time"] = time.Now()
		condition["login_ip"] = ctx.RemoteIP()
		user.Update(condition)

		builder := strings.Builder{}
		builder.WriteString("user:info:")
		builder.WriteString(strconv.FormatUint(uint64(user.Id), 10))
		key := builder.String()

		cache.Del(key)
		token, tokenErr := tool.JwtToken(user.Id)
		data["token"] = token
		if nil != tokenErr {
			response.Error(ctx, err.ERROR, err.TOKEN_ERROR)

			return
		}

		response.Success(ctx, data)

		return
	}

	response.Error(ctx, err.RUNTIME_ERROR, err.PASSWORD_VALIDATE_ERROR)
}

func Register(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	rePassword := ctx.PostForm("re_password")

	if password != rePassword {
		response.Error(ctx, err.RUNTIME_ERROR, err.PASSWORD_DIFF_ERROR)

		return
	}

	userModel := &model.User{}

	condition := make(map[string]interface{})
	condition["user_name"] = userName

	_, exists := userModel.Get(condition)

	if exists {
		response.Error(ctx, err.RUNTIME_ERROR, err.USER_NAME_EXISTS_ERROR)

		return
	}

	sale := tool.UUID()
	encryptPassword := tool.Password(password, sale)

	newUser := &model.User{
		Nick:          "",
		UserName:      userName,
		Password:      encryptPassword,
		Salt:          sale,
		LastLoginTime: time.Now(),
		LoginIp:       ctx.RemoteIP(),
	}

	id, success := userModel.Create(newUser)

	if success {
		token, tokenErr := tool.JwtToken(id)

		if nil != tokenErr {
			response.Error(ctx, err.ERROR, err.TOKEN_ERROR)

			return
		}
		successMap := make(map[string]interface{})
		successMap["token"] = token

		response.Success(ctx, successMap)

		return
	}

	response.Error(ctx, err.RUNTIME_ERROR, err.USER_CREATE_ERROR)
}
