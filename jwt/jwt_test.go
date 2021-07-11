package jwt

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	token := NewToken().SetOwner("Korbin").String()
	fmt.Println(token)
	fmt.Println(Check(token))
	fmt.Println(Check(token + "1"))
	fmt.Println(IsExpire(token))
	fmt.Printf("%+v\n", GetPayload(token))
	token = Refresh(token)
	fmt.Println(token)
	fmt.Printf("%+v\n", GetPayload(token))
}
