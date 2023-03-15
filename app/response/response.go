package response

import (
	"encoding/json"
	"github.com/TskFok/GinApi/app/err"
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/utils/kafka"
	"github.com/TskFok/GinApi/utils/logger"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, list interface{}) {
	successInfo := make(map[string]interface{})
	successInfo["code"] = err.SUCCESS
	successInfo["msg"] = err.GetMsg(err.SUCCESS)
	successInfo["data"] = list

	infoLog(ctx, successInfo)

	ctx.JSON(err.SUCCESS, successInfo)
}

func Error(ctx *gin.Context, code int, info any) {
	errorInfo := make(map[string]interface{})
	errorInfo["data"] = make(map[string]interface{})

	switch info.(type) {
	case int:
		errorInfo["code"] = info.(int)
		errorInfo["msg"] = err.GetMsg(info.(int))
	case string:
		errorInfo["code"] = err.RuntimeError
		errorInfo["msg"] = info.(string)
	}

	infoLog(ctx, errorInfo)

	ctx.JSON(code, errorInfo)
}

func infoLog(ctx *gin.Context, info interface{}) {
	if global.Env == "debug" {
		buf, jsonErr := json.Marshal(info)

		if nil != jsonErr {
			logger.Error(jsonErr.Error())
		}

		logger.Info(string(buf))

		return
	}

	//获取请求头
	header := make(map[string]interface{})
	for i, v := range ctx.Request.Header {
		header[i] = v[0]
	}

	//获取请求内容
	body := make(map[string]interface{})
	for i, v := range ctx.Request.PostForm {
		body[i] = v[0]
	}

	//获取请求内容
	rawQuery := ctx.Request.URL.RawQuery

	res := make(map[string]interface{})
	res["response"] = info
	res["request"] = body
	res["header"] = header
	res["query"] = rawQuery

	buf, jsonErr := json.Marshal(res)

	if nil != jsonErr {
		logger.Error(jsonErr.Error())
	}

	//发送至kafka
	kafka.Send(string(buf), "log")
}
