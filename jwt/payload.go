package jwt

import (
	"errors"
	"time"
)

const (
	defaultPayloadIssuer  = "jwt"
	defaultPayloadPurpose = "authentication"

	RecipientBrowser = "browser"
	RecipientAndroid = "android"
	RecipientIOS     = "ios"
	RecipientWechat  = "wechat mini program"
)

type Payload struct {
	// 签发者
	Issuer string `json:"issuer,omitempty"`
	// 令牌所有者,存放ID等标识
	Owner string `json:"owner,omitempty"`
	// 用途,默认值authentication表示用于登录认证
	Purpose string `json:"purpose,omitempty"`
	// 接受方,表示申请该令牌的设备来源,如浏览器、Android等
	Recipient string `json:"recipient,omitempty"`
	// 令牌签发时间
	Time time.Time `json:"time"`
	// 过期时间, expire = time + duration
	Expire time.Time `json:"expire,omitempty"`
	// 令牌持续时间,即生命周期,0则表示没有设置过期时间
	Duration time.Duration `json:"duration,omitempty"`
	// 其他扩展的自定义参数
	External External `json:"external,omitempty"`
}

func defaultPayload() *Payload {
	return &Payload{
		Issuer:   defaultPayloadIssuer,
		Purpose:  defaultPayloadPurpose,
		Time:     time.Now(),
		External: make(External, 0),
	}
}

// Valid Verify that the field of Payload are valid
// 校验Payload的属性是否有效
func (payload *Payload) Valid() error {
	if payload.Time.IsZero() {
		return errors.New("invalid time of payload, can not be nil")
	}
	if payload.Duration != 0 && payload.Expire.Before(payload.Time) {
		return errors.New("invalid expire of payload, can not before time")
	}
	return nil
}
