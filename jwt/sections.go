package jwt

import "time"

const (
	defaultHeaderAlgorithm   = "HSA256"
	defaultHeaderType        = "JWT"
	defaultPayloadIssuer     = "whereabouts.icu"
	defaultPayloadOwner      = ""
	defaultPayloadPurpose    = "Authentication"
	defaultPayloadRecipient  = "Browser"
	defaultPayloadExpireTime = time.Minute * 30
	defaultSignature         = "HeZebin"
)

type Header struct {
	Algorithm string `json:"algorithm"`
	Type      string `json:"type"`
}

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
	Time int64 `json:"time"`
	// 过期时间
	Expire int64 `json:"expire"`
	// 令牌持续时间,即生命周期
	Duration time.Duration `json:"duration"`
	// 其他扩展的自定义参数
	External map[string]interface{} `json:"external"`
}

func defaultHeader() *Header {
	return &Header{Algorithm: defaultHeaderAlgorithm, Type: defaultHeaderType}
}

func defaultPayload() *Payload {
	currentTime := time.Now()
	return &Payload{
		Issuer:    defaultPayloadIssuer,
		Owner:     defaultPayloadOwner,
		Purpose:   defaultPayloadPurpose,
		Recipient: defaultPayloadRecipient,
		Time:      currentTime.Unix(),
		Expire:    currentTime.Add(defaultPayloadExpireTime).Unix(),
		Duration:  defaultPayloadExpireTime,
		External:  nil,
	}
}

var (
	signature = defaultSignature
)

func SetSignature(sign string) {
	signature = sign
}

func getSignature() string {
	return signature
}
