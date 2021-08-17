package jwt

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/jwt/alg"
	"github.com/whereabouts/sdk/utils/stringer"
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
	signatureJson, err := token.Algorithm.Encrypt(SplicingSegment(header, payload), secret)
	if err != nil {
		return "", errors.Wrap(err, "encrypt signature err")
	}
	signature := EncodeSegment(signatureJson)
	token.Signature = signature
	if err != nil {
		return "", err
	}
	token.Raw = SplicingSegment(header, payload, signature)
	return token.Raw, nil
}

func (token *Token) String(secret string) (string, error) {
	return token.Sign(secret)
}

// Valid Verify that the token is valid by comparing the signature generated after
// the header and payload are encrypted with the same algorithm and key with the original signature.
// 校验token是否有效, 通过将header和payload以同样的算法和密钥加密后生成的signature与原signature对比实现
func (token *Token) Valid(secret string) bool {
	if stringer.IsEmpty(token.Signature) {
		return false
	}
	segments := Split2Segment(token.Raw)
	if len(segments) != 3 {
		return false
	}
	signatureJson, err := token.Algorithm.Encrypt(SplicingSegment(segments[0], segments[1]), secret)
	if err != nil {
		return false
	}
	return token.Signature == EncodeSegment(signatureJson)
}

func (token *Token) Expired() bool {
	return time.Now().After(token.Payload.Expire)
}

// Refresh Reset the expiration time based on the current time according to the duration of the token
// 在当前时间基础上根据token的持续时间重置过期时间
func (token *Token) Refresh() *Token {
	token.Payload.Expire = time.Now().Add(token.Payload.Duration)
	return token
}

func Parse(token string) (*Token, error) {
	segments := Split2Segment(token)
	if len(segments) != 3 {
		return nil, errors.New("a token consists of three segments connected by points")
	}

	kernel := &Token{Raw: token}
	// parse header
	headerJson, err := DecodeSegment(segments[0])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(headerJson, &kernel.Header)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal header err")
	}
	algorithm := alg.GetAlg(kernel.Header["alg"])
	if algorithm == nil {
		return nil, errors.New("algorithm does not exist in alg.manager")
	}
	kernel.Algorithm = algorithm
	// parse payload
	payloadJson, err := DecodeSegment(segments[1])
	err = json.Unmarshal(payloadJson, kernel.Payload)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal payload err")
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

// SetDuration Set the token validity time, the expiration time will be reset based on the current time.
// 设置token有效时间, 将在当前时间基础上重置过期时间
func (token *Token) SetDuration(duration time.Duration) *Token {
	token.Payload.Duration = duration
	token.SetExpire(time.Now().Add(duration))
	return token
}

func (token *Token) SetExternal(external External) *Token {
	token.Payload.External = external
	return token
}
