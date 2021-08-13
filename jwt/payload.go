package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

const (
	defaultPayloadIssuer     = "jwt"
	defaultPayloadOwner      = ""
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
	External map[string]interface{} `json:"external"`
}

func NewPayload(owner string) *Payload {
	return &Payload{
		Owner:    owner,
		Time:     time.Now(),
		External: make(map[string]interface{}, 0),
	}
}

func DefaultPayload() *Payload {
	currentTime := time.Now()
	return &Payload{
		Issuer:    defaultPayloadIssuer,
		Owner:     defaultPayloadOwner,
		Purpose:   defaultPayloadPurpose,
		Recipient: defaultPayloadRecipient,
		Time:      currentTime,
		Expire:    currentTime.Add(defaultPayloadExpireTime),
		Duration:  defaultPayloadExpireTime,
		External:  make(map[string]interface{}, 0),
	}
}

func (payload *Payload) EncodeBase64URL() (string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "marshal payload err")
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

func (payload *Payload) SetIssuer(issuer string) *Payload {
	payload.Issuer = issuer
	return payload
}

func (payload *Payload) SetOwner(owner string) *Payload {
	payload.Owner = owner
	return payload
}

func (payload *Payload) SetPurpose(purpose string) *Payload {
	payload.Purpose = purpose
	return payload
}

func (payload *Payload) SetRecipient(recipient string) *Payload {
	payload.Recipient = recipient
	return payload
}

func (payload *Payload) SetTime(time time.Time) *Payload {
	payload.Time = time
	return payload
}

func (payload *Payload) SetExpire(expire time.Time) *Payload {
	payload.Expire = expire
	return payload
}

func (payload *Payload) SetDuration(duration time.Duration) *Payload {
	payload.Duration = duration
	return payload
}

func (payload *Payload) SetExternal(external map[string]interface{}) *Payload {
	payload.External = external
	return payload
}
