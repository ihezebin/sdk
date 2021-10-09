package aliyun

import (
	"encoding/json"
)

type Message struct {
	// SignName 短信签名名称，eg: "阿里云"
	SignName string `mapstructure:"sign_name" json:"sign_name"`
	// TemplateCode 短信模板CODE
	TemplateCode string `mapstructure:"template_code" json:"template_code"`
	// TemplateParam 短信模板变量对应的实际值，eg：{"code":"1234"}
	TemplateParamJson string `mapstructure:"template_param_json" json:"template_param_json"`
}

func NewMessage() *Message {
	return &Message{}
}

func (msg *Message) WithTemplate(code string, param map[string]interface{}) *Message {
	paramJson, _ := json.Marshal(param)
	msg.TemplateCode = code
	msg.TemplateParamJson = string(paramJson)
	return msg
}

func (msg *Message) WithSignName(sign string) *Message {
	msg.SignName = sign
	return msg
}
