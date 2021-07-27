package test

import (
	"fmt"
	"github.com/whereabouts/sdk/storage/redis"
	"log"
	"testing"
)

func TestFunc(t *testing.T) {
	client, err := redis.Init(redis.Config{
		Addr:      ":6379",
		Password:  "root",
		MaxIdle:   10,
		MaxActive: 50,
	})
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	cache := redis.New("email")
	val, err := cache.Get("12113").Value()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("val: ", val)
}
