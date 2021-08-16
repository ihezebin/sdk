package jwt

import (
	"encoding/base64"
	"strings"
)

// EncodeSegment encode JWT specific base64url encoding with padding stripped
func EncodeSegment(seg []byte) string {
	//return strings.TrimRight(base64.URLEncoding.EncodeToString(seg), "=")
	return base64.URLEncoding.EncodeToString(seg)
}

// DecodeSegment decode JWT specific base64url encoding with padding stripped
func DecodeSegment(seg string) ([]byte, error) {
	//if l := len(seg) % 4; l > 0 {
	//	seg += strings.Repeat("=", 4-l)
	//}

	return base64.URLEncoding.DecodeString(seg)
}

func splicingSegment(a, b string) string {
	return strings.Join([]string{a, b}, ".")
}

func splitSegment(v string) []string {
	return strings.Split(v, ".")
}
