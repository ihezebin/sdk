package code

import (
	"fmt"
	"math/rand"
	"strings"
)

func Default() string {
	return Generate(DefaultLen, TypeBlend)
}

func String(length int, codeType int) string {
	return Generate(length, codeType)
}

func Blend(length int) string {
	return Generate(length, TypeBlend)
}

func Digit(length int) string {
	return Generate(length, TypeDigit)
}

func LetterLower(length int) string {
	return Generate(length, TypeLetterLower)
}

func LetterUpper(length int) string {
	return Generate(length, TypeLetterUpper)
}

func Generate(length int, codeType int) (code string) {
	for i := 0; i < length; i++ {
		switch codeType {
		case TypeBlend:
			code = fmt.Sprintf("%s%c", code, characterBlend[rand.Intn(len(characterBlend))])
		case TypeDigit:
			code = fmt.Sprintf("%s%c", code, characterDigit[rand.Intn(len(characterDigit))])
		case TypeLetterLower:
			code = fmt.Sprintf("%s%c", code, characterLetterLower[rand.Intn(len(characterLetterLower))])
		case TypeLetterUpper:
			code = fmt.Sprintf("%s%c", code, characterLetterUpper[rand.Intn(len(characterLetterUpper))])
		default:
			code = fmt.Sprintf("%s%c", code, characterBlend[rand.Intn(len(characterBlend))])
		}
	}
	return code
}

func Verify(in string, code string, capslock ...bool) bool {
	if len(capslock) == 0 || !capslock[0] {
		return strings.ToUpper(in) == strings.ToUpper(code)
	} else {
		return in == code
	}
}
