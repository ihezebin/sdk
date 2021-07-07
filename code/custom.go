package code

import (
	"fmt"
	"math/rand"
)

func Custom(template string) string {
	return NewCustom(DefaultLen, template)
}

func NewCustom(length int, template string) (code string) {
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
