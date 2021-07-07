package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

func encodeStruct2Base64URL(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	code := base64.URLEncoding.EncodeToString(data)
	return code
}

func encodeString2Base64URL(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func decodeBase64URL2String(s string) (string, error) {
	code, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(code), nil
}

func decodeBase64URL2Struct(s string, obj interface{}) error {
	code, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	err = json.Unmarshal(code, obj)
	if err != nil {
		return err
	}
	return nil
}

func encryptionSHA256(s string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
