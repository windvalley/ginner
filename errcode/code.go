package errcode

import "net/http"

var (
	OK = &ErrCode{
		Status:  http.StatusOK,
		Code:    "OK",
		Message: "OK",
	}

	InternalServerError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "InternalError",
		Message: "Internal server error",
	}

	ServerPanicError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "PanicError",
		Message: "System meets up panic error, please contact admin",
	}

	MysqlDBError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "MysqlError",
		Message: "MySQL error",
	}

	// How to use:
	// err := errcode.New(errcode.ErrDataNotExist, nil)
	// errFormat := err.Add("somestring")
	// handler.SendResponse(c, errFormat, nil)
	DataNotExistError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "NoData",
		Message: "%s has no data.",
	}

	ValidationError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "ValidationError",
		Message: "Validation failed.",
	}

	// user/client errors

	ArgsNotFoundError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "ParameterMissing",
		Message: "The parameter %s was not found",
	}

	ArgsValueError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "ArgsValueError",
		Message: "The value of %s was invalid",
	}

	PostDataError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "PostDataError",
		Message: "The post data was invalid.",
	}

	RecordNotFoundError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "RecordNotFound",
		Message: "record not found in database.",
	}

	UserNotFoundError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "UserNotFound",
		Message: "The user was not found.",
	}

	PasswordIncorrectError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "PasswordInvalid",
		Message: "The password was incorrect.",
	}

	TokenInvalidError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "TokenInvalid",
		Message: "The token was invalid.",
	}
)
