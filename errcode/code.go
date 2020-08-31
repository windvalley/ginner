package errcode

import "net/http"

var (
	// OK success response
	OK = &ErrCode{
		Status:  http.StatusOK,
		Code:    "OK",
		Message: "OK",
	}

	// follows are error responses of server-side

	// InternalServerError unknown errors of inside server
	InternalServerError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "InternalError",
		Message: "Internal server error",
	}

	// ServerPanicError panic error
	ServerPanicError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "PanicError",
		Message: "Server meets up panic error, please contact admin.",
	}

	// DBError Usage:
	// err := errcode.New(errcode.ErrDataNotExist, nil)
	// err1 := err.Add("somestring")
	// handler.SendResponse(c, err1, nil)
	DBError = &ErrCode{
		Status:  http.StatusInternalServerError,
		Code:    "DBError",
		Message: "%v",
	}

	// follows are error responses of client-side

	ValidationError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "ValidationError",
		Message: "%v",
	}

	UserNotFoundError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "UserNotFound",
		Message: "The user is not found.",
	}

	PasswordIncorrectError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "PasswordIncorrect",
		Message: "The password is incorrect.",
	}

	TokenInvalidError = &ErrCode{
		Status:  http.StatusUnauthorized,
		Code:    "TokenInvalid",
		Message: "%v",
	}

	APISignError = &ErrCode{
		Status:  http.StatusUnauthorized,
		Code:    "SignatureInvalid",
		Message: "%v",
	}

	AccessForbiddenError = &ErrCode{
		Status:  http.StatusForbidden,
		Code:    "AccessForbidden",
		Message: "%s",
	}

	RecordNotFoundError = &ErrCode{
		Status:  http.StatusBadRequest,
		Code:    "RecordNotFound",
		Message: "Record not found.",
	}

	TooManyRequestError = &ErrCode{
		Status:  http.StatusTooManyRequests,
		Code:    "TooManyRequest",
		Message: "too many request.",
	}
)
