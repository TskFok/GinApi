package err

var MsgFlags = map[int]string{
	SUCCESS:                 "成功",
	ERROR:                   "系统错误",
	RUNTIME_ERROR:           "程序运行错误",
	UNDEFINED_ERROR:         "信息不存在",
	PASSWORD_DIFF_ERROR:     "二次密码不相同",
	USER_NAME_EXISTS_ERROR:  "用户名已存在",
	UNKNWON_ERROR:           "未知错误",
	USER_UNDEFINED_ERROR:    "用户不存在",
	PASSWORD_VALIDATE_ERROR: "密码错误",
	USER_CREATE_ERROR:       "用户创建失败",
}

func getMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

func GetSuccess(data interface{}) map[string]interface{} {
	successInfo := make(map[string]interface{})
	successInfo["code"] = SUCCESS
	successInfo["msg"] = getMsg(SUCCESS)
	successInfo["data"] = data

	return successInfo
}

func GetErrorInfo(code int) map[string]interface{} {
	errorInfo := make(map[string]interface{})
	errorInfo["code"] = code
	errorInfo["msg"] = getMsg(code)
	errorInfo["data"] = make(map[string]interface{})

	return errorInfo
}
