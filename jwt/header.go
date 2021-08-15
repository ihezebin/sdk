package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
)

const (
	defaultHeaderType = "JWT"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func (h *header) EncodeBase64URL() (string, error) {
	data, err := json.Marshal(h)
	if err != nil {
		return "", errors.Wrap(err, "marshal header err")
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

func (h *header) Algorithm() string {
	return h.Alg
}

func (h *header) Type() string {
	return h.Typ
}
