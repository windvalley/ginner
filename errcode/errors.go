package errcode

import "errors"

var (
	// ErrRecordNotFound internal use
	ErrRecordNotFound = errors.New("record not found")

	// ErrRecordExists internal use
	ErrRecordExists = errors.New("record already exists")
)
