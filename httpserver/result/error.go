package result

import "encoding/json"

type Error struct {
	HttpStatusCode int         `json:"-"`
	Code           interface{} `json:"code"`
	Message        string      `json:"message"`
	Data           interface{} `json:"data,omitempty"`
}

func (err *Error) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}

func (err *Error) WithHttpStatusCode(httpStatusCode int) *Error {
	err.HttpStatusCode = httpStatusCode
	return err
}

func NewError(code interface{}, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

func Err2Error(err error, code interface{}) *Error {
	if err == nil {
		return nil
	}
	// do not need convert
	if e, ok := err.(*Error); ok {
		return e
	}
	return Error(code, err.Error())
}
