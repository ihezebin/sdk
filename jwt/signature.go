package jwt

import (
	"encoding/base64"
	"github.com/whereabouts/sdk/jwt/algorithm"
)

type signature struct {
	algorithm algorithm.Algorithm
	secret    string
	data      []byte
}

func (s *signature) EncodeBase64URL() (string, error) {
	return base64.URLEncoding.EncodeToString(s.data), nil
}

func (s *signature) Algorithm() algorithm.Algorithm {
	return s.algorithm
}

func (s *signature) Secret() string {
	return s.secret
}
