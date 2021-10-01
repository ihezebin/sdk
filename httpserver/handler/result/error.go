package result

import (
	"github.com/pkg/errors"
	"net/http"
)

type Err struct {
	statusCode int
	Code       interface{} `json:"code"`
	Message    string      `json:"message"`
}

func (err *Err) Error() string {
	return err.Message
}

func (err *Err) WithStatusCode(statusCode int) *Err {
	err.statusCode = statusCode
	return err
}

func (err *Err) StatusCode() int {
	if err.statusCode == 0 {
		return http.StatusOK
	}
	return err.statusCode
}

func Error(code interface{}, msg string) *Err {
	return &Err{
		Code:    code,
		Message: msg,
	}
}

func WrapError(err error, code interface{}, msg string) *Err {
	return Error2Err(errors.Wrap(err, msg), code)
}

func WrapErrorf(err error, code interface{}, format string, args ...interface{}) *Err {
	return Error2Err(errors.Wrapf(err, format, args...), code)
}

func Error2Err(err error, code interface{}) *Err {
	if err == nil {
		return nil
	}
	// do not need convert
	if e, ok := err.(*Err); ok {
		return e
	}
	return Error(code, err.Error())
}
