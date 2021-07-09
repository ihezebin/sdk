package http_error

type JSON map[string]interface{}

func DefaultJSON(code interface{}, message string) interface{} {
	return struct {
		Code    interface{} `json:"code"`
		Message string      `json:"message"`
	}{code, message}
}
