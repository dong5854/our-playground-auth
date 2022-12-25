package customerror

import (
	"errors"
	"fmt"
)

const (
	ErrNotDefined = iota + 1
	ErrInternalServer
	ErrBadRequest
	ErrDBConnection
	ErrDBInternal
)

type Error struct {
	code   int
	msg    string
	origin error
}

func (e *Error) Error() string {
	if e.origin != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.origin)
	}

	return e.msg
}

func (e *Error) Unwrap() error {
	return e.origin
}

func (e *Error) Code() int {
	return e.code
}

func Wrap(origin error, code int, format string, args ...interface{}) error {
	return &Error{
		code:   code,
		msg:    fmt.Sprintf(format, args...),
		origin: origin,
	}
}

func New(code int, format string, args ...interface{}) error {
	return Wrap(nil, code, format, args...)
}

func GetCode(err error) int {
	var customErr *Error
	if !errors.As(err, &customErr) {
		return ErrNotDefined
	}

	return customErr.Code()
}
