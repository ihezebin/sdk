package result

type JSON map[string]interface{}

type DefaultResp struct {
	Code    interface{} `json:"coder"`
	Message string      `json:"message"`
}

func DefaultJSON(code interface{}, message string) interface{} {
	return struct {
		Code    interface{} `json:"coder"`
		Message string      `json:"message"`
	}{code, message}
}
