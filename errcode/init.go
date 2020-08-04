package errcode

import "fmt"

type ErrCode struct {
	Code, Message string
}

func (ec ErrCode) Error() string {
	return ec.Message
}

type Err struct {
	ErrCode
	Err error
}

// Err is also error type
func (e *Err) Error() string {
	return fmt.Sprintf("Err - code: %s, message: %s, error: %s",
		e.Code, e.Message, e.Err)
}

// replace origin Message field's %s or %d and so on
func (e *Err) Add(item ...interface{}) error {
	e.Message = fmt.Sprintf(e.Message, item...)
	return e
}

// add additional content for Message field.
func (e *Err) Addf(format string, a ...interface{}) error {
	e.Message += " " + fmt.Sprintf(format, a...)
	return e
}

func New(ec *ErrCode, err error) *Err {
	return &Err{
		ErrCode: ErrCode{
			Code:    ec.Code,
			Message: ec.Message,
		},
		Err: err,
	}
}

// get Code and Message field from ErrCode、Err、error
func DecodeErr(err error) (string, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *ErrCode:
		return typed.Code, typed.Message
	}

	return InternalServerError.Code, err.Error()
}
