package jwt

import (
	"github.com/whereabouts/sdk/jwt/alg"
	"github.com/whereabouts/sdk/logger"
	"testing"
)

func TestJWT(t *testing.T) {
	token, err := New(alg.HSA256()).SetOwner("jwt").Sign("sign")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(token)
}
