package code

import (
	"fmt"
	"math/rand"
	"strings"
)

func Default() string {
	return Generate(DefaultLen, TypeBlend)
}

func String(length int, codeType CodeType) string {
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

func Generate(length int, codeType CodeType) (code string) {
	for i := 0; i < length; i++ {
		switch codeType {
		case TypeBlend:
			code = fmt.Sprintf("%s%c", code, TemplateBlend[rand.Intn(len(TemplateBlend))])
		case TypeDigit:
			code = fmt.Sprintf("%s%c", code, TemplateDigit[rand.Intn(len(TemplateDigit))])
		case TypeLetterLower:
			code = fmt.Sprintf("%s%c", code, TemplateLetterLower[rand.Intn(len(TemplateLetterLower))])
		case TypeLetterUpper:
			code = fmt.Sprintf("%s%c", code, TemplateLetterUpper[rand.Intn(len(TemplateLetterUpper))])
		default:
			code = fmt.Sprintf("%s%c", code, TemplateBlend[rand.Intn(len(TemplateBlend))])
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
