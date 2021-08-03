package result

import "encoding/json"

type HttpError struct {
	HttpStatusCode int         `json:"-"`
	Code           interface{} `json:"coder"`
	Message        string      `json:"message"`
}

func (err *HttpError) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}

func (err *HttpError) WithHttpStatusCode(httpStatusCode int) *HttpError {
	err.HttpStatusCode = httpStatusCode
	return err
}

func Error(code interface{}, msg string) *HttpError {
	return &HttpError{
		Code:    code,
		Message: msg,
	}
}

func Err2HttpError(err error, code interface{}) *HttpError {
	if err == nil {
		return nil
	}
	// do not need convert
	if e, ok := err.(*HttpError); ok {
		return e
	}
	return Error(code, err.Error())
}
