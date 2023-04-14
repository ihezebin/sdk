package result

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type Err struct {
	statusCode int
	Message    string `json:"message"`
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

func Error(msg string) *Err {
	return &Err{
		Message: msg,
	}
}

func Errorf(format string, args ...interface{}) *Err {
	return &Err{
		Message: fmt.Sprintf(format, args...),
	}
}

func WrapError(err error, msg string) *Err {
	return Error2Err(errors.Wrap(err, msg))
}

func WrapErrorf(err error, code interface{}, format string, args ...interface{}) *Err {
	return Error2Err(errors.Wrapf(err, format, args...))
}

func Error2Err(err error) *Err {
	if err == nil {
		return nil
	}
	// do not need convert
	if e, ok := err.(*Err); ok {
		return e
	}
	return Error(err.Error())
}
