package result

import (
	"fmt"
	"net/http"
)

type Result struct {
	statusCode int
	Code       bool        `json:"code"`
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

func Failed(msg string) *Result {
	return &Result{Code: false, Message: msg}
}

func FailedWithFormat(format string, args ...interface{}) *Result {
	return &Result{Code: false, Message: fmt.Sprintf(format, args...)}
}

func FailedWithErr(err error) *Result {
	return &Result{Code: false, Message: err.Error()}
}

func (res *Result) WithSubResult(subResult *SubResult) *Result {
	res.Data = subResult
	return res
}

type SubResult struct {
	SubCode    interface{} `json:"sub_code"`
	SubData    interface{} `json:"sub_data"`
	SubMessage string      `json:"sub_message"`
}

func NewSubResult(subCode interface{}, subMsg string) *SubResult {
	return &SubResult{
		SubCode:    subCode,
		SubMessage: subMsg,
	}
}

func (sub *SubResult) WithMessage(subMsg string) *SubResult {
	sub.SubMessage = subMsg
	return sub
}

func (sub *SubResult) WithCode(subCode interface{}) *SubResult {
	sub.SubCode = subCode
	return sub
}

func (sub *SubResult) WithData(subData interface{}) *SubResult {
	sub.SubData = subData
	return sub
}
