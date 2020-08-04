package errcode

var (
	OK = &ErrCode{Code: "OK", Message: "OK"}

	InternalServerError = &ErrCode{
		Code:    "InternalError",
		Message: "Internal server error.",
	}

	ErrMysqlDB = &ErrCode{
		Code:    "MysqlError",
		Message: "MySQL error.",
	}

	// How to use:
	// err := errcode.New(errcode.ErrDataNotExist, nil)
	// errFormat := err.Add("somestring")
	// handler.SendResponse(c, errFormat, nil)
	ErrDataNotExist = &ErrCode{
		Code:    "NoData",
		Message: "%s has no data.",
	}

	ErrValidation = &ErrCode{
		Code:    "ValidationError",
		Message: "Validation failed.",
	}

	// user/client errors
	ErrArgsNotFound = &ErrCode{
		Code: "ParameterMissing",
		Message: "The input parameter 'timestart' or 'timeend' that " +
			"is mandatory for processing this request was not supplied " +
			"or its value was null.",
	}

	ErrPostData = &ErrCode{
		Code:    "PostDataError",
		Message: "The post data was invalid.",
	}

	ErrRecordNotFound = &ErrCode{
		Code:    "RecordNotFound",
		Message: "record not found in database.",
	}

	ErrUserNotFound = &ErrCode{
		Code:    "UserNotFound",
		Message: "The user was not found.",
	}

	ErrPasswordIncorrect = &ErrCode{
		Code:    "PasswordInvalid",
		Message: "The password was incorrect.",
	}

	ErrTokenInvalid = &ErrCode{
		Code:    "TokenInvalid",
		Message: "The token was invalid.",
	}
)
