package tool

import "github.com/TskFok/GinApi/app/err"

func GetSuccess(data interface{}) map[string]interface{} {
	successInfo := make(map[string]interface{})
	successInfo["code"] = err.SUCCESS
	successInfo["msg"] = err.GetMsg(err.SUCCESS)
	successInfo["data"] = data

	return successInfo
}

func GetErrorInfo(code int) map[string]interface{} {
	errorInfo := make(map[string]interface{})
	errorInfo["code"] = code
	errorInfo["msg"] = err.GetMsg(code)
	errorInfo["data"] = make(map[string]interface{})

	return errorInfo
}
