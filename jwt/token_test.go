package jwt

import (
	"fmt"
	"github.com/whereabouts/sdk/jwt/alg"
	"github.com/whereabouts/sdk/logger"
	"testing"
	"time"
)

const secret = "jwt.whereabouts.icu"

func TestJWT(t *testing.T) {
	token := New(alg.HSA256()).SetOwner("jwt").SetDuration(time.Second*2).Claim("name", "Korbin")
	sign, err := token.Sign(secret)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(sign)
	fmt.Println("valid:", token.Valid(secret))
	time.Sleep(time.Second * 5)
	fmt.Println("expired:", token.Expired())
	fmt.Println("====refresh====")
	sign, err = token.Refresh().Sign(secret)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(sign)
	fmt.Println("valid:", token.Valid(secret))
	fmt.Println("expired:", token.Expired())
}
