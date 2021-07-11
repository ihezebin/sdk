package code

import (
	"fmt"
)

func PrintCNChar() {
	for _, v := range TemplateZnCn {
		fmt.Printf("%c", v)
	}
}

func ZnCn() string {
	return NewCustom(DefaultLen, TemplateZnCn)
}

func NewZnCn(length int) string {
	return NewCustom(length, TemplateZnCn)
}
