package result

import "net/http"

type Result struct {
	statusCode int
	Code       interface{} `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (res *Result) StatusCode() int {
	if res.statusCode == 0 {
		return http.StatusOK
	}
	return res.statusCode
}

func (res *Result) WithStatusCode(statusCode int) *Result {
	res.statusCode = statusCode
	return res
}

func (res *Result) WithCode(code interface{}) *Result {
	if code == nil {
		res.Code = true
	} else {
		res.Code = code
	}
	return res
}

func (res *Result) WithMessage(msg string) *Result {
	res.Message = msg
	return res
}

func New() *Result {
	return &Result{}
}

func Succeed(data interface{}) *Result {
	return &Result{Code: true, Data: data}
}

func Failed(err error) *Result {
	return &Result{Code: false, Message: err.Error()}
}
