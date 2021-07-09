package http_error

import "encoding/json"

const (
	CodeBoolSuccess bool = true
	CodeBoolOk      bool = true
	CodeBoolFail    bool = false
	//CodeBasic    int = 1000
	//CodeResource int = 2000
	//CodeAuth     int = 3000
	//CodeNull     int = 4000
	//CodeSystem   int = 5000
	//CodeParam    int = 6000
	//CodeConvert  int = 7000
)

type HttpError struct {
	HttpStatusCode int         `json:"-"`
	ErrCode        interface{} `json:"err_code"`
	ErrMessage     string      `json:"err_message"`
	Code           interface{} `json:"code"`
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
		Code:       code,
		Message:    msg,
		ErrCode:    code,
		ErrMessage: msg,
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
