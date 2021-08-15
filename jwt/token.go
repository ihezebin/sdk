package jwt

import (
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/jwt/algorithm"
	"strings"
	"time"
)

type Token struct {
	header    *header
	payload   *payload
	signature *signature
}

func New(algorithm algorithm.Algorithm) *Token {
	return NewWithClaims(algorithm)
}

func NewWithClaims(algorithm algorithm.Algorithm, claims ...Claim) *Token {
	return &Token{
		&header{
			Typ: defaultHeaderType,
			Alg: algorithm.String(),
		},
		newPayload(claims...),
		&signature{algorithm: algorithm},
	}
}

func (token *Token) Claim(key string, value interface{}) *Token {
	token.payload.Claim(key, value)
	return token
}

func (token *Token) Sign(secret string) (string, error) {
	h, err := token.Header().EncodeBase64URL()
	if err != nil {
		return "", err
	}
	p, err := token.Payload().EncodeBase64URL()
	if err != nil {
		return "", err
	}
	token.Signature().secret = secret
	token.Signature().data, err = token.Signature().Algorithm().Encrypt(token)
	if err != nil {
		return "", errors.Wrap(err, "encrypt signature err")
	}
	s, err := token.Signature().EncodeBase64URL()
	if err != nil {
		return "", err
	}
	return strings.Join([]string{h, p, s}, "."), nil
}

func (token *Token) String(secret string) (string, error) {
	return token.Sign(secret)
}

func (token *Token) Payload() *payload {
	return token.payload
}

func (token *Token) Header() *header {
	return token.header
}

func (token *Token) Signature() *signature {
	return token.signature
}

func (token *Token) SetIssuer(issuer string) *Token {
	token.payload.Issuer = issuer
	return token
}

func (token *Token) SetOwner(owner string) *Token {
	token.payload.Owner = owner
	return token
}

func (token *Token) SetPurpose(purpose string) *Token {
	token.payload.Purpose = purpose
	return token
}

func (token *Token) SetRecipient(recipient string) *Token {
	token.payload.Recipient = recipient
	return token
}

func (token *Token) SetExpire(expire time.Duration) *Token {
	token.payload.Expire = token.payload.Time.Add(expire)
	return token
}

func (token *Token) SetDuration(duration time.Duration) *Token {
	token.payload.Duration = duration
	token.SetExpire(duration)
	return token
}

func (token *Token) SetExternal(external map[string]interface{}) *Token {
	token.payload.External = external
	return token
}

//func Check(token string) bool {
//	sections := strings.Split(token, ".")
//	if len(sections) != 3 {
//		return false
//	}
//	sign := encryptionSHA256(fmt.Sprintf("%s.%s", sections[0], sections[1]), getSignature())
//	return sections[2] == sign
//}
//
//func IsExpire(token string) bool {
//	payload := GetPayload(token)
//	if time.Now().Unix() > payload.Expire {
//		return true
//	}
//	return false
//}
//
//func Refresh(token string) string {
//	p := GetPayload(token)
//	p.Expire = time.Now().Add(p.Duration).Unix()
//	header := encodeStruct2Base64URL(defaultHeader())
//	payload := encodeStruct2Base64URL(p)
//	sign := encryptionSHA256(fmt.Sprintf("%s.%s", header, payload), getSignature())
//	return fmt.Sprintf("%s.%s.%s", header, payload, sign)
//}
//
//func Parse(token string) Payload {
//	payload := Payload{}
//	sections := strings.Split(token, ".")
//	if len(sections) != 3 {
//		return payload
//	}
//	_ = decodeBase64URL2Struct(sections[1], &payload)
//	return payload
//}
