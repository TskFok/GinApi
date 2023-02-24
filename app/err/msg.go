package err

var MsgFlags = map[int]string{
	SUCCESS:               "成功",
	ERROR:                 "系统错误",
	RUNTIME_ERROR:         "程序运行错误",
	SECOND_PASSWORD_ERROR: "二次密码不相同",
	USER_NAME_EXISTS:      "用户名已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
