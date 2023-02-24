package user

import (
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/model"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(ctx *gin.Context) {
	//userName := ctx.PostForm("user_name")
	//password := ctx.PostForm("password")

}

func Register(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	rePassword := ctx.PostForm("re_password")

	emptyMap := make(map[string]interface{})

	if password != rePassword {
		ctx.JSON(err.SUCCESS, gin.H{
			"code": err.SECOND_PASSWORD_ERROR,
			"msg":  err.GetMsg(err.SECOND_PASSWORD_ERROR),
			"data": emptyMap,
		})

		return
	}

	userModel := &model.User{}

	condition := make(map[string]interface{})
	condition["user_name"] = userName

	user := userModel.HasOneByName(condition)

	if user {
		ctx.JSON(err.SUCCESS, gin.H{
			"code": err.USER_NAME_EXISTS,
			"msg":  err.GetMsg(err.USER_NAME_EXISTS),
			"data": emptyMap,
		})

		return
	}

	newUser := &model.User{
		Nick:          "aaa",
		UserName:      userName,
		Password:      password,
		Salt:          "aaa",
		LastLoginTime: time.Now(),
		LoginIp:       "aaa",
	}

	res := userModel.CreateUser(newUser)

	if res {
		ctx.JSON(err.SUCCESS, gin.H{
			"code": err.SUCCESS,
			"mgs":  err.GetMsg(err.SUCCESS),
			"data": emptyMap,
		})

		return
	}

	ctx.JSON(err.RUNTIME_ERROR, gin.H{
		"code": err.RUNTIME_ERROR,
		"mgs":  err.GetMsg(err.RUNTIME_ERROR),
		"data": emptyMap,
	})
}
