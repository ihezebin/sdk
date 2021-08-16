package alg

import (
	"crypto/hmac"
	"crypto/sha256"
)

type hsa256 struct {
}

func HSA256() *hsa256 {
	return &hsa256{}
}

func init() {
	RegisterAlg(HSA256())
}

func (alg *hsa256) Name() string {
	return "HSA256"
}

func (alg *hsa256) Encrypt(signing, secret string) ([]byte, error) {
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(signing))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
