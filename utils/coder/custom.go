package coder

import (
	"fmt"
	"math/rand"
)

func Custom(template Template) string {
	return NewCustom(DefaultLen, template)
}

func NewCustom(length int, template Template) (code string) {
	for i := 0; i < length; i++ {
		index := rand.Intn(len(template) / 3)
		for j, v := range template {
			if j == index*3 {
				code = fmt.Sprintf("%s%c", code, v)
				break
			}
		}
	}
	return code
}
