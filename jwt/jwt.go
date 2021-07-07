package jwt

import (
	"fmt"
	"strings"
	"time"
)

type JWT struct {
	Header    *Header  `json:"header"`
	Payload   *Payload `json:"payload"`
	Signature string   `json:"signature"`
}

func NewToken() *JWT {
	return &JWT{
		Header:    defaultHeader(),
		Payload:   defaultPayload(),
		Signature: getSignature(),
	}
}

func Check(token string) bool {
	sections := strings.Split(token, ".")
	if len(sections) != 3 {
		return false
	}
	sign := encryptionSHA256(fmt.Sprintf("%s.%s", sections[0], sections[1]), getSignature())
	return sections[2] == sign
}

func IsExpire(token string) bool {
	payload := GetPayload(token)
	if time.Now().Unix() > payload.Expire {
		return true
	}
	return false
}

func Refresh(token string) string {
	p := GetPayload(token)
	p.Expire = time.Now().Add(p.Duration).Unix()
	header := encodeStruct2Base64URL(defaultHeader())
	payload := encodeStruct2Base64URL(p)
	sign := encryptionSHA256(fmt.Sprintf("%s.%s", header, payload), getSignature())
	return fmt.Sprintf("%s.%s.%s", header, payload, sign)
}

func GetPayload(token string) Payload {
	payload := Payload{}
	sections := strings.Split(token, ".")
	if len(sections) != 3 {
		return payload
	}
	_ = decodeBase64URL2Struct(sections[1], &payload)
	return payload
}

func (jwt *JWT) String() string {
	header := encodeStruct2Base64URL(jwt.Header)
	payload := encodeStruct2Base64URL(jwt.Payload)
	sign := encryptionSHA256(fmt.Sprintf("%s.%s", header, payload), jwt.Signature)
	return fmt.Sprintf("%s.%s.%s", header, payload, sign)
}

func (jwt *JWT) GetPayload() Payload {
	return *jwt.Payload
}

func (jwt *JWT) SetIssuer(issuer string) *JWT {
	jwt.Payload.Issuer = issuer
	return jwt
}

func (jwt *JWT) SetOwner(owner string) *JWT {
	jwt.Payload.Owner = owner
	return jwt
}

func (jwt *JWT) SetPurpose(purpose string) *JWT {
	jwt.Payload.Purpose = purpose
	return jwt
}

func (jwt *JWT) SetRecipient(recipient string) *JWT {
	jwt.Payload.Recipient = recipient
	return jwt
}

func (jwt *JWT) SetExpire(expire time.Duration) *JWT {
	jwt.Payload.Expire = time.Unix(jwt.Payload.Time, 0).Add(expire).Unix()
	return jwt
}

func (jwt *JWT) SetDuration(duration time.Duration) *JWT {
	jwt.Payload.Duration = duration
	jwt.SetExpire(duration)
	return jwt
}

func (jwt *JWT) SetExternal(external map[string]interface{}) *JWT {
	jwt.Payload.External = external
	return jwt
}
