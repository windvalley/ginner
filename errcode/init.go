package errcode

import (
	"fmt"
)

// ErrCode wrap up response messages
type ErrCode struct {
	Status        int
	Code, Message string
}

func (ec ErrCode) Error() string {
	return ec.Message
}

// Err wrap up system errors
type Err struct {
	ErrCode
	SysErr error
}

// New wrap up the err which will not show to client, but write in error.log
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

// Error Err is also error type
func (e *Err) Error() string {
	return fmt.Sprintf("Err - status: %d, code: %s, message: %s, error: %s",
		e.Status, e.Code, e.Message, e.SysErr)
}

// Add replace origin Message field's %s or %d and so on
func (e *Err) Add(item ...interface{}) {
	e.Message = fmt.Sprintf(e.Message, item...)
}

// Addf add additional content for Message field
func (e *Err) Addf(format string, a ...interface{}) {
	e.Message += " " + fmt.Sprintf(format, a...)
}

// DecodeErr get Status, Code and Message field from ErrCode、Err、error for showing to client,
// and get SysErr field from Err for logging in the server local log file.
func DecodeErr(err error) (int, string, string, string) {
	if err == nil {
		return OK.Status, OK.Code, OK.Message, ""
	}

	switch v := err.(type) {
	case *Err:
		if v.SysErr != nil {
			return v.Status, v.Code, v.Message, v.SysErr.Error()
		}
		return v.Status, v.Code, v.Message, ""
	case *ErrCode:
		return v.Status, v.Code, v.Message, ""
	default:
		return InternalServerError.Status, InternalServerError.Code, err.Error(), ""
	}
}
