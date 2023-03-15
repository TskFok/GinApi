package err

const (
	Success = 200
	Error   = 500

	RuntimeError          = 400
	UnknownError          = 400000
	PasswordValidateError = 400001
	PasswordDiffError     = 400002
	UserNameExistsError   = 400003
	UserCreateError       = 400004
	RedisError            = 400005
	RouteCreateError      = 400006
	RouteUpdateError      = 400007

	UndefinedError       = 404
	ParamsUndefinedError = 404000
	UserUndefinedError   = 404001
	RouterUndefinedError = 404002

	TokenError = 500000
)
