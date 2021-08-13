package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/jwt/algorithm"
)

const (
	defaultHeaderTypeJWT = "JWT"
)

type Header struct {
	Algorithm algorithm.Algorithm
	Type      string
}

func NewHeader(alg algorithm.Algorithm, typ string) *Header {
	return &Header{alg, typ}
}

func DefaultHeader() *Header {
	return &Header{Algorithm: algorithm.HSA256(), Type: defaultHeaderTypeJWT}
}

func (header *Header) EncodeBase64URL() (string, error) {
	data, err := json.Marshal(map[string]string{
		"algorithm": header.Algorithm.String(),
		"type":      header.Type,
	})
	if err != nil {
		return "", errors.Wrap(err, "marshal header err")
	}
	return base64.URLEncoding.EncodeToString(data), nil
}
