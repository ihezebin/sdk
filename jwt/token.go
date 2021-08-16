package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/jwt/alg"
	"github.com/whereabouts/sdk/utils/stringer"
	"strings"
	"time"
)

type Token struct {
	Raw       string
	Header    map[string]string
	Payload   *Payload
	Signature string
	Algorithm alg.Algorithm
}

func Default() *Token {
	return NewWithClaims(alg.HSA256())
}

func New(algorithm alg.Algorithm) *Token {
	return NewWithClaims(algorithm)
}

func NewWithClaims(algorithm alg.Algorithm, claims ...Claim) *Token {
	alg.RegisterAlg(algorithm)
	return &Token{
		Header: map[string]string{
			"typ": "JWT",
			"alg": algorithm.Name(),
		},
		Payload:   newPayload(claims...),
		Algorithm: algorithm,
	}
}

func (token *Token) Sign(secret string) (string, error) {
	headerJson, err := json.Marshal(token.Header)
	if err != nil {
		return "", err
	}
	payloadJson, err := json.Marshal(token.Payload)
	if err != nil {
		return "", err
	}
	header := EncodeSegment(headerJson)
	payload := EncodeSegment(payloadJson)
	signatureJson, err := token.Algorithm.Encrypt(splicingSegment(header, payload), secret)
	if err != nil {
		return "", errors.Wrap(err, "encrypt signature err")
	}
	signature := EncodeSegment(signatureJson)
	token.Signature = signature
	if err != nil {
		return "", err
	}
	token.Raw = strings.Join([]string{header, payload, signature}, ".")
	return token.Raw, nil
}

func (token *Token) String(secret string) (string, error) {
	return token.Sign(secret)
}

func (token *Token) Valid(secret string) bool {
	if stringer.IsEmpty(token.Signature) {
		return false
	}
	segments := splitSegment(token.Raw)
	if len(segments) != 3 {
		return false
	}
	signatureJson, err := token.Algorithm.Encrypt(splicingSegment(segments[0], segments[1]), secret)
	if err != nil {
		return false
	}
	fmt.Println(segments[0])
	fmt.Println(segments[1])
	fmt.Println(segments[2])
	fmt.Println(token.Signature)
	fmt.Println(EncodeSegment(signatureJson))
	return token.Signature == EncodeSegment(signatureJson)
}

func (token *Token) Expired(secret string) bool {
	return true
}

func (token *Token) Refresh(secret string) *Token {
	return nil
}

func Parse(token string) (*Token, error) {
	segments := splitSegment(token)
	if len(segments) != 3 {
		return nil, errors.New("")
	}

	kernel := &Token{Raw: token}
	// parse header
	headerJson, err := DecodeSegment(segments[0])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(headerJson, &kernel.Header)
	if err != nil {
		return nil, errors.Wrap(err, "parse header err")
	}
	algorithm := alg.GetAlg(kernel.Header["alg"])
	if algorithm == nil {
		return nil, errors.New("alg err")
	}
	kernel.Algorithm = algorithm
	// parse payload
	payloadJson, err := DecodeSegment(segments[1])
	err = json.Unmarshal(payloadJson, kernel.Payload)
	if err != nil {
		return nil, errors.Wrap(err, "parse payload err")
	}
	// parse signature
	kernel.Signature = segments[2]

	return kernel, nil
}

func (token *Token) Claim(key string, value interface{}) *Token {
	token.Payload.External[key] = value
	return token
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

func (token *Token) SetExpire(expire time.Time) *Token {
	token.Payload.Expire = expire
	return token
}

// SetDuration Set the token validity time, the issuance time will be reset to the current time, and the expiration time is the duration after now.
// 设置token有效时间, 将重置签发时间为当前时间, 过期时间为此后的duration
func (token *Token) SetDuration(duration time.Duration) *Token {
	token.Payload.Time = time.Now()
	token.Payload.Duration = duration
	token.SetExpire(token.Payload.Time.Add(duration))
	return token
}

func (token *Token) SetExternal(external External) *Token {
	token.Payload.External = external
	return token
}
