package code

import (
	"fmt"
)

func PrintCNChar() {
	for _, v := range zn_cn {
		fmt.Printf("%c", v)
	}
}

func ZN_CN() string {
	return NewCustom(DefaultLen, zn_cn)
}

func NewZN_CN(length int) (code string) {
	return NewCustom(length, zn_cn)
}
