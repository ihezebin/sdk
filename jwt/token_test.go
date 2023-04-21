package jwt

import (
	"fmt"
	"github.com/ihezebin/sdk/jwt/alg"
	"github.com/ihezebin/sdk/logger"
	"github.com/ihezebin/sdk/utils/timer"
	"testing"
	"time"
)

const secret = "jwt.whereabouts.icu"

const fakeToken = "eyJhbGciOiJIU0EyNTYiLCJ0eXAiOiJKV1QifQ" +
	".eyJpc3N1ZXIiOiJqd3QiLCJvd25lciI6Imp3dCIsInB1cnBvc2UiOiJhdXRoZW50aWNhdGlvbiIsInRpbWUiOiIyMDIxLTA4LTIyVDA0OjA0OjQ2LjQxNjg0NjIrMDg6MDAiLCJleHBpcmUiOiIyMDIxLTA4LTIyVDA0OjA0OjQ4LjQxNjg0NjIrMDg6MDAiLCJkdXJhdGlvbiI6MjAwMDAwMDAwMCwiZXh0ZXJuYWwiOnsibmFtZSI6IktvcmJpbiJ9fQ" +
	".fdbeAeDc5Hq0abzhU76FWXT_XVS2H9ryGgPvew1WSNU1"

func TestJWT(t *testing.T) {
	token := New(alg.HSA256()).SetOwner("username/id").SetDuration(time.Second*2).Claim("name", "Korbin")
	sign, err := token.Sign(secret)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(sign)
	fmt.Println(timer.Format(token.Payload.Time))
	fmt.Println(timer.Format(token.Payload.Expire))
	fmt.Println(token.Payload.External["name"])
	fmt.Println(token.Payload.Owner)
	fmt.Println(token.Payload.Issuer)
	fmt.Println(token.Payload.Purpose)
}

func TestNoExpired(t *testing.T) {
	jwt := Default().SetOwner("hezebin")
	sign, err := jwt.Sign(secret)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(sign)
	fmt.Println(jwt.Valid(secret) == nil)
	fmt.Println(jwt.Expired())
}

func TestValidAndParse(t *testing.T) {
	token := NewWithClaims(
		alg.HSA256(),
		WithExpire(time.Now().Add(5*time.Second)),
		WithCustom("name", "Korbin"),
	)
	sign, err := token.Sign(secret)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(sign)
	fmt.Println(token.Valid(secret) == nil)
	// fake token
	fake, err := Parse(fakeToken)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(fake.Valid(secret) == nil)
}

func TestRefreshAndExpired(t *testing.T) {
	tokenStr, err := NewWithClaims(
		alg.HSA256(),
		WithExpire(time.Now().Add(2*time.Second)),
		WithCustom("name", "Korbin"),
	).String(secret)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(tokenStr)
	token, err := Parse(tokenStr)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(token.Expired())
	time.Sleep(time.Second * 4)
	fmt.Println(token.Expired())

	// Refresh
	tokenStr, err = token.Refresh(secret)
	if err != nil {
		logger.Fatal(err)
	}
	token, err = Parse(tokenStr)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(token.Expired())
	fmt.Println(token.Sign(secret))
}
