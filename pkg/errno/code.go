package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	Error               = &Errno{Code: 10003, Message: "系统繁忙"}

	ErrValidation = &Errno{Code: 20001, Message: "参数填写有误"}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrTokenInvalid       = &Errno{Code: 20101, Message: "The token was invalid."}
	ErrAuthInvalid        = &Errno{Code: 20102, Message: "权限无效"}
	ErrUserNameOrPassword = &Errno{Code: 20103, Message: "用户名或密码错误"}
	UserFreeze            = &Errno{Code: 20104, Message: "用户账号已被冻结"}
	//permission errors
	//ERR            = &Errno{Code: 20201, Message: "用户账号已被冻结"}
)
