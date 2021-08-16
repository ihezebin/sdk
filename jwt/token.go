package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/jwt/alg"
	"strings"
	"time"
)

type Token struct {
	Header    map[string]interface{}
	Payload   *Payload
	Signature string
	Algorithm alg.Algorithm
}

func New(algorithm alg.Algorithm) *Token {
	return NewWithClaims(algorithm)
}

func NewWithClaims(algorithm alg.Algorithm, claims ...Claim) *Token {
	return &Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": algorithm.String(),
		},
		Payload: newPayload(claims...),
		Algorithm: algorithm,
	}
}

func (token *Token) Claim(key string, value interface{}) *Token {
	token.Payload.External[key] = value
	return token
}

func (token *Token) Sign(secret string) (string, error) {
	header, err := encodeSegment(token.Header)
	if err != nil {
		return "", err
	}
	payload, err := encodeSegment(token.Payload)
	if err != nil {
		return "", err
	}
	sign, err = token.Algorithm.Encrypt(header, payload, secret)
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


func (token *Token) SetIssuer(issuer string) *Token {
	token.Payload.Issuer = issuer
	return token
}

func (token *Token) SetOwner(owner string) *Token {
	token.Payload.Owner = owner
	return token
}

func (token *Token) SetPurpose(purpose string) *Token {
	token.Payload.Purpose = purpose
	return token
}

func (token *Token) SetRecipient(recipient string) *Token {
	token.Payload.Recipient = recipient
	return token
}

func (token *Token) SetExpire(expire time.Duration) *Token {
	token.Payload.Expire = token.Payload.Time.Add(expire)
	return token
}

func (token *Token) SetDuration(duration time.Duration) *Token {
	token.Payload.Duration = duration
	token.SetExpire(duration)
	return token
}

func (token *Token) SetExternal(external map[string]interface{}) *Token {
	token.Payload.External = external
	return token
}

func encodeSegment(segment interface{}) (string, error) {
	switch segment.(type) {
	case []byte:
		return base64.URLEncoding.EncodeToString(segment.([])), nil
	}
	data, err := json.Marshal(segment)
	if err != nil {
		return "", errors.Wrap(err, "marshal token segment err")
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

func decodeSegment(data string, segment interface{}) error {
	decodeData, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return errors.Wrap(err, "decode token segment err")
	}
	err = json.Unmarshal(decodeData, segment)
	if err != nil {
		return errors.Wrap(err, "unmarshal token segment err")
	}
	return nil
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
