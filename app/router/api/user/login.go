package user

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/TskFok/GinApi/app/tool"
	"github.com/gin-gonic/gin"
	"time"
)

func Info(ctx *gin.Context) {
	if user, exists := ctx.Get("user"); exists {
		ctx.JSON(err.SUCCESS, err.GetSuccess(user))
		return
	}

	ctx.JSON(err.UNDEFINED_ERROR, err.GetErrorInfo(err.USER_UNDEFINED_ERROR))
}

func Login(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")

	userModel := &model.User{}

	condition := make(map[string]interface{})
	condition["user_name"] = userName

	user, exists := userModel.HasOneByName(condition)

	if !exists {
		ctx.JSON(err.UNDEFINED_ERROR, err.GetErrorInfo(err.USER_UNDEFINED_ERROR))

		return
	}

	if tool.Password(password, user.Salt) == user.Password {
		data := make(map[string]interface{})
		data["token"] = tool.JwtToken(user.Id)
		ctx.JSON(err.SUCCESS, err.GetSuccess(data))

		return
	}

	ctx.JSON(err.RUNTIME_ERROR, err.GetErrorInfo(err.PASSWORD_VALIDATE_ERROR))
}

func Register(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	rePassword := ctx.PostForm("re_password")

	if password != rePassword {
		ctx.JSON(err.RUNTIME_ERROR, err.GetErrorInfo(err.PASSWORD_DIFF_ERROR))

		return
	}

	userModel := &model.User{}

	condition := make(map[string]interface{})
	condition["user_name"] = userName

	_, exists := userModel.HasOneByName(condition)

	if exists {
		ctx.JSON(err.RUNTIME_ERROR, err.GetErrorInfo(err.USER_NAME_EXISTS_ERROR))

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

	id, success := userModel.CreateUser(newUser)

	if success {
		token := tool.JwtToken(id)
		successMap := make(map[string]interface{})
		successMap["token"] = token

		ctx.JSON(err.SUCCESS, err.GetSuccess(successMap))

		return
	}

	ctx.JSON(err.RUNTIME_ERROR, err.GetErrorInfo(err.USER_CREATE_ERROR))
}
