package err

import "net/http"

var MsgFlags = map[int]string{
	http.StatusOK:                  "成功",
	http.StatusInternalServerError: "系统错误",
	http.StatusBadRequest:          "程序运行错误",
	http.StatusNotFound:            "信息不存在",
	PasswordDiffError:              "二次密码不相同",
	UserNameExistsError:            "用户名已存在",
	UnknownError:                   "未知错误",
	ParamsUndefinedError:           "字段不存在",
	ParamsValidateError:            "字段验证出错",
	UserUndefinedError:             "用户不存在",
	RouterUndefinedError:           "路由不存在",
	PasswordValidateError:          "密码错误",
	UserCreateError:                "用户创建失败",
	RedisError:                     "redis错误",
	RouteHasExistsError:            "router已存在",
	RouteNotExistsError:            "router不存在",
	RouteCreateError:               "router创建失败",
	RouteUpdateError:               "router更新失败",
	TokenError:                     "token获取失败",
	RouteRepeatError:               "router重复",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[http.StatusInternalServerError]
}
