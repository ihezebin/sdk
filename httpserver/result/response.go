package result

type Response struct {
	HttpStatusCode int         `json:"-"`
	Code           interface{} `json:"code"`
	Message        string      `json:"message"`
	Data           interface{} `json:"data,omitempty"`
}

func (resp *Response) WithStatusCode() *Response {

}

func ResponseSuccess(data interface{}) *Response {
	return &Response{Code: true, Data: data}
}

func ResponseError(err error) *Response {
	return &Response{Code: true, Message: err.Error()}
}
