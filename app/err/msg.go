package err

var MsgFlags = map[int]string{
	SUCCESS:                 "成功",
	ERROR:                   "系统错误",
	RUNTIME_ERROR:           "程序运行错误",
	UNDEFINED_ERROR:         "信息不存在",
	PASSWORD_DIFF_ERROR:     "二次密码不相同",
	USER_NAME_EXISTS_ERROR:  "用户名已存在",
	UNKNWON_ERROR:           "未知错误",
	PARAMS_UNDEFINED_ERROR:  "字段不存在",
	USER_UNDEFINED_ERROR:    "用户不存在",
	ROUTER_UNDEFINED_ERROR:  "路由不存在",
	PASSWORD_VALIDATE_ERROR: "密码错误",
	USER_CREATE_ERROR:       "用户创建失败",
	REDIS_ERROR:             "redis错误",
	ROUTE_CREATE_ERROR:      "router创建失败",
	ROUTE_UPDATE_ERROR:      "router更新失败",
	TOKEN_ERROR:             "token获取失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
