package alg

import (
	"crypto/hmac"
	"crypto/sha256"
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

func (alg *hsa256) Encrypt(header, payload, secret string) ([]byte, error) {
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(strings.Join([]string{header, payload}, ".")))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
