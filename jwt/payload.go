package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

const (
	defaultPayloadIssuer     = "jwt"
	defaultPayloadPurpose    = "authentication"
	defaultPayloadRecipient  = "browser"
	defaultPayloadExpireTime = time.Minute * 30
)

type payload struct {
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

func NewPayload() *payload {
	return defaultPayload()
}

func defaultPayload() *payload {
	currentTime := time.Now()
	return &payload{
		Issuer:    defaultPayloadIssuer,
		Purpose:   defaultPayloadPurpose,
		Recipient: defaultPayloadRecipient,
		Time:      currentTime,
		Expire:    currentTime.Add(defaultPayloadExpireTime),
		Duration:  defaultPayloadExpireTime,
		External:  make(map[string]interface{}, 0),
	}
}

func (p *payload) EncodeBase64URL() (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", errors.Wrap(err, "marshal payload err")
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

func (p *payload) SetIssuer(issuer string) *payload {
	p.Issuer = issuer
	return p
}

func (p *payload) SetOwner(owner string) *payload {
	p.Owner = owner
	return p
}

func (p *payload) SetPurpose(purpose string) *payload {
	p.Purpose = purpose
	return p
}

func (p *payload) SetRecipient(recipient string) *payload {
	p.Recipient = recipient
	return p
}

func (p *payload) SetTime(time time.Time) *payload {
	p.Time = time
	return p
}

func (p *payload) SetExpire(expire time.Time) *payload {
	p.Expire = expire
	return p
}

func (p *payload) SetDuration(duration time.Duration) *payload {
	p.Duration = duration
	return p
}

func (p *payload) SetExternal(external map[string]interface{}) *payload {
	p.External = external
	return p
}

func (p *payload) Claim(key string, value interface{}) *payload {
	p.External[key] = value
	return p
}
