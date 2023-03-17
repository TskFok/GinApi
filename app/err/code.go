package err

const (
	UnknownError          = 400000
	PasswordValidateError = 400001
	PasswordDiffError     = 400002
	UserNameExistsError   = 400003
	UserCreateError       = 400004
	RedisError            = 400005
	RouteCreateError      = 400006
	RouteUpdateError      = 400007
	ParamsValidateError   = 400008
	RouteHasExistsError   = 400009
	RouteNotExistsError   = 400010
	RouteRepeatError      = 400011

	ParamsUndefinedError = 404000
	UserUndefinedError   = 404001
	RouterUndefinedError = 404002

	TokenError = 500000
)
