package errorx

import (
	"errors"
	"fmt"

	"github.com/oxidnova/novadm/backend/driver/schema/code"
)

// Error is an custom error with code.
type Error struct {
	Code code.Code
	error

	a Additional
}

type Additional map[string]any

func (a Additional) Get(k string) interface{} {
	v, _ := a[k]
	return v
}

// WrapError returns an error annotating err with a code at the point
// Wrap is called, and the supplied reason message.
// If err is nil, Wrap returns nil.
func WrapError(code code.Code, err error) *Error {
	if err == nil {
		return nil
	}

	return &Error{
		error: err,
		Code:  code,
	}
}

// Errorf returns an error annotating err with a code at the point Wrapf is called,
// and the format specifier. If err is nil, Wrapf returns nil.
func Errorf(code code.Code, format string, args ...interface{}) *Error {
	return &Error{
		error: fmt.Errorf(format, args...),
		Code:  code,
	}
}

// ConvertError convert an error to Error
func ConvertError(err error) *Error {
	var target *Error
	if errors.As(err, &target) {
		return target
	}

	return WrapError(code.Unknown, err)
}

var (
	ErrUnauthorized = &Error{Code: code.Unauthorized, error: errors.New("unauthorized")}
)

func (e *Error) Additional(k string, v interface{}) *Error {
	if e == nil {
		return e
	}

	if e.a == nil {
		e.a = Additional{}
	}

	e.a[k] = v
	return e
}

func (e *Error) GetAdditional() Additional {
	if e.a == nil {
		e.a = Additional{}
	}
	return e.a
}

func (e *Error) Get(key string) interface{} {
	if e == nil || e.a == nil {
		return nil
	}

	return e.a.Get(key)
}

func (e *Error) GetInt(key string) int {
	v := e.Get(key)

	switch vv := v.(type) {
	case int:
		return vv
	case int32:
		return int(vv)
	case int64:
		return int(vv)
	case uint:
		return int(vv)
	case uint32:
		return int(vv)
	case uint64:
		return int(vv)
	case float32:
		return int(vv)
	case float64:
		return int(vv)
	}

	return 0
}

func (e *Error) GetString(key string) string {
	v := e.Get(key)
	s, _ := v.(string)
	return s
}
