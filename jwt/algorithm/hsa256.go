package algorithm

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/whereabouts/sdk/jwt"
	"strings"
)

func HSA256() Algorithm {
	return Algorithm{
		Name:    "HSA256",
		Handler: SHA256Handler,
	}
}

func SHA256Handler(header *jwt.Header, payload *jwt.Payload, secret string) (string, error) {
	headerBase64URL, err := header.EncodeBase64URL()
	if err != nil {
		return "", err
	}
	payloadBase64URL, err := payload.EncodeBase64URL()
	if err != nil {
		return "", err
	}
	h := hmac.New(sha256.New, []byte(secret))
	_, err = h.Write([]byte(strings.Join([]string{headerBase64URL, payloadBase64URL}, ".")))
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(h.Sum(nil)), nil
}
