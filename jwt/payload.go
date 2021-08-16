package jwt

import (
	"time"
)

const (
	defaultPayloadIssuer     = "jwt"
	defaultPayloadPurpose    = "authentication"
	defaultPayloadRecipient  = "browser"
	defaultPayloadExpireTime = time.Minute * 30
)

type Payload struct {
	// 签发者
	Issuer string `json:"issuer"`
	// 令牌所有者,存放ID等标识
	Owner string `json:"owner"`
	// 用途,默认值authentication表示用于登录认证
	Purpose string `json:"purpose"`
	// 接受方,表示申请该令牌的设备来源,如浏览器、Android等
	Recipient string `json:"recipient"`
	// 令牌签发时间
	Time time.Time `json:"time"`
	// 过期时间
	Expire time.Time `json:"expire"`
	// 令牌持续时间,即生命周期
	Duration time.Duration `json:"duration"`
	// 其他扩展的自定义参数
	External External `json:"external"`
}

func defaultPayload() *Payload {
	currentTime := time.Now()
	return &Payload{
		Issuer:    defaultPayloadIssuer,
		Purpose:   defaultPayloadPurpose,
		Recipient: defaultPayloadRecipient,
		Time:      currentTime,
		Expire:    currentTime.Add(defaultPayloadExpireTime),
		Duration:  defaultPayloadExpireTime,
		External:  make(map[string]interface{}, 0),
	}
}

func (payload *Payload) Claim(key string, value interface{}) *Payload {
	payload.External[key] = value
	return payload
}
