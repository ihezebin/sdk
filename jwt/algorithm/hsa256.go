package algorithm

import (
	"crypto/hmac"
	"crypto/sha256"
	"github.com/whereabouts/sdk/jwt"
	"strings"
)

type hsa256 struct {
}

func HSA256() *hsa256 {
	return &hsa256{}
}

func (alg *hsa256) String() string {
	return "HSA256"
}

func (alg *hsa256) Encrypt(token *jwt.Token) ([]byte, error) {
	header, err := token.Header().EncodeBase64URL()
	if err != nil {
		return nil, err
	}
	payload, err := token.Payload().EncodeBase64URL()
	if err != nil {
		return nil, err
	}
	h := hmac.New(sha256.New, []byte(token.Signature().Secret()))
	_, err = h.Write([]byte(strings.Join([]string{header, payload}, ".")))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
