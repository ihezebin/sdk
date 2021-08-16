package jwt

import (
	"fmt"
	"github.com/whereabouts/sdk/jwt/alg"
	"github.com/whereabouts/sdk/logger"
	"testing"
)

const secret = "jwt.whereabouts.icu"

func TestJWT(t *testing.T) {
	token := New(alg.HSA256()).SetOwner("jwt")
	sign, err := token.Sign(secret)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(sign)
	fmt.Println(token.Valid(secret))
}
