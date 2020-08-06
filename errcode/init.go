package errcode

import (
	"fmt"
	"net/http"
)

type ErrCode struct {
	Status        int
	Code, Message string
}

func (ec ErrCode) Error() string {
	return ec.Message
}

type Err struct {
	ErrCode
	SysErr error
}

// wrap up the err which will not show to client, but write in error.log
func New(ec *ErrCode, err error) *Err {
	return &Err{
		ErrCode: ErrCode{
			Status:  ec.Status,
			Code:    ec.Code,
			Message: ec.Message,
		},
		SysErr: err,
	}
}

// Err is also error type
func (e *Err) Error() string {
	return fmt.Sprintf("Err - status: %d, code: %s, message: %s, error: %s",
		e.Status, e.Code, e.Message, e.SysErr)
}

// replace origin Message field's %s or %d and so on
func (e *Err) Add(item ...interface{}) error {
	e.Message = fmt.Sprintf(e.Message, item...)
	return e
}

// add additional content for Message field
func (e *Err) Addf(format string, a ...interface{}) error {
	e.Message += " " + fmt.Sprintf(format, a...)
	return e
}

// get Status, Code and Message field from ErrCode、Err、error for showing to client
func DecodeErr(err error) (int, string, string) {
	if err == nil {
		return OK.Status, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Status, typed.Code, typed.Message
	case *ErrCode:
		return typed.Status, typed.Code, typed.Message
	default:
		return http.StatusInternalServerError, InternalServerError.Code, err.Error()
	}
}
