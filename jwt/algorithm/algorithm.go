package algorithm

import (
	"github.com/whereabouts/sdk/jwt"
)

type Algorithm interface {
	String() string
	Encrypt(token *jwt.Token) ([]byte, error)
}
