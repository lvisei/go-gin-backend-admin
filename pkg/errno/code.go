package errno

var (
	OK = &Errno{Code: 20000, Message: "OK"}
	// Common errors
	INVALID_PARAMS = &Errno{Code: 40000, Message: "Request parameter error."}

	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrDatabase         = &Errno{Code: 10002, Message: "Database error."}
	ErrBind             = &Errno{Code: 10003, Message: "Error occurred while binding the request body to the struct."}

	ERROR_AUTH                     = &Errno{Code: 20001, Message: "Token error."}
	ERROR_AUTH_TOKEN               = &Errno{Code: 20002, Message: "Token generation failed."}
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = &Errno{Code: 20003, Message: "Token has timed out."}
	ERROR_AUTH_CHECK_TOKEN_FAIL    = &Errno{Code: 20004, Message: "Token authentication failed."}
	ERROR_Not_TOKEN_EXIST          = &Errno{Code: 20005, Message: "The length of the `Authorization` header is zero.."}

	// user errors
	ERROR_Encrypt            = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ERROR_User_Not_Found     = &Errno{Code: 20102, Message: "The user was not found."}
	ERROR_EXIST_USER         = &Errno{Code: 20103, Message: "The user was exist."}
	ErrTokenInvalid          = &Errno{Code: 20104, Message: "The token was invalid."}
	ERROR_Password_Incorrect = &Errno{Code: 20105, Message: "The password was incorrect."}
	ERROR_ADD_User_FAIL      = &Errno{Code: 20106, Message: "Failed to Create user."}
)
