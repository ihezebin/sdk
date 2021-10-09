package tencent

import (
	"strconv"
)

type Message struct {
	//SmsSdkAppId 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666
	SmsSdkAppId string `mapstructure:"sms_sdk_app_id" json:"sms_sdk_app_id"`
	// SignName 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台] 查看
	SignName string `mapstructure:"sign_name" json:"sign_name"`
	// TemplateId 模板 ID: 必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台] 查
	TemplateId string `mapstructure:"template_id" json:"template_id"`
	// TemplateParamSet 模板参数: 若无模板参数，则设置为空
	TemplateParamSet []string `mapstructure:"template_param_set" json:"template_param_set"`
	// ExtendCode 短信码号扩展号: 默认未开通，如需开通请联系 [sms helper]
	ExtendCode string `mapstructure:"extend_code" json:"extend_code"`
	// SessionContext 用户的 session 内容: 可以携带用户侧 ID 等上下文信息，server 会原样返回
	SessionContext string `mapstructure:"session_context" json:"session_context"`
	// SenderId 国际/港澳台短信 SenderId: 国内短信填空，默认未开通，如需开通请联系 [sms helper]
	SenderId string `mapstructure:"sender_id" json:"sender_id"`
}

func NewMessage() *Message {
	return &Message{}
}

func (msg *Message) WithTemplate(id string, params ...int) *Message {
	paramSet := make([]string, 0, len(params))
	for _, param := range params {
		paramSet = append(paramSet, strconv.Itoa(param))
	}
	msg.TemplateId = id
	msg.TemplateParamSet = paramSet
	return msg
}

func (msg *Message) WithAppId(appId string) *Message {
	msg.SmsSdkAppId = appId
	return msg
}

func (msg *Message) WithSignName(sign string) *Message {
	msg.SignName = sign
	return msg
}

// 不常用参数

func (msg *Message) WithExtendCode(extendCode string) *Message {
	msg.ExtendCode = extendCode
	return msg
}

func (msg *Message) WithSessionContext(sessionContext string) *Message {
	msg.SessionContext = sessionContext
	return msg
}

func (msg *Message) WithSenderId(senderId string) *Message {
	msg.SenderId = senderId
	return msg
}
