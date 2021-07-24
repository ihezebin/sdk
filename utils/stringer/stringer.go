package stringer

import "strings"

func Length(s string) int {
	return len(s)
}

func IsEmpty(s string) bool {
	return Length(s) == 0 || s == ""
}

func NotEmpty(s string) bool {
	return !IsEmpty(s)
}

func CharAt(s string, index int) string {
	return string([]rune(s)[index])
}

func Equals(a string, b string) bool {
	return strings.Compare(a, b) == 0
}

func EqualsIgnoreCase(a string, b string) bool {
	return strings.EqualFold(a, b)
}

func GetBytes(s string) []byte {
	return []byte(s)
}

func StartWithPrefix(s string, prefix string) bool {
	return StartWithPrefixOffset(s, prefix, 0)
}

func StartWithPrefixOffset(s string, prefix string, offset int) bool {
	return Equals(s[offset:len(prefix)], prefix)
}

func EndWithSuffix(s string, suffix string) bool {
	return Equals(s[len(s)-len(suffix):], suffix)
}

func SubstringWithBegin(s string, begin int) string {
	return SubstringWithRange(s, begin, len(s))
}

func SubstringWithRange(s string, begin int, end int) string {
	return s[begin:end]
}

func Contact(a string, b string) string {
	return string(append([]byte(a), []byte(b)...))
}
