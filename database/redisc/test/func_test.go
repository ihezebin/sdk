package test

import (
	"fmt"
	"github.com/whereabouts/sdk/database/redisc"
	"github.com/whereabouts/sdk/logger"
	"testing"
)

func TestFunc(t *testing.T) {
	client, err := redisc.Init(redisc.Config{
		Addr:      ":6379",
		Password:  "root",
		MaxIdle:   10,
		MaxActive: 50,
	})
	defer client.Close()
	if err != nil {
		logger.Fatal(err)
	}
	cache := redisc.New("email")
	val, err := cache.Get("12113").Value()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("val: ", val)
}
