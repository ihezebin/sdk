package test

import (
	"fmt"
	"github.com/whereabouts/sdk/database/redisc"
	"github.com/whereabouts/sdk/utils/coder"
	"log"
	"testing"
)

func TestRedis(t *testing.T) {
	client, err := redisc.Init(redisc.Config{
		Addr:      ":6379",
		Password:  "root",
		MaxIdle:   10,
		MaxActive: 50,
	})
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	id := "korbin"
	result := GetEmailCache().AddEmailCode(id, coder.Default())
	fmt.Println(result)
	result = GetEmailCache().GetEmailCode(id)
	fmt.Println(result.String())
}
