package result

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Err struct {
	HttpStatusCode int         `json:"-"`
	Code           interface{} `json:"code"`
	Message        string      `json:"message"`
}

func (err *Err) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}

func (err *Err) WithStatusCode(statusCode int) *Err {
	err.HttpStatusCode = statusCode
	return err
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
