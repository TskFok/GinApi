package err

var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "系统错误",
	INVALID_PARAMS: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
