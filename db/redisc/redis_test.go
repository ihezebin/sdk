package redisc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb, err := NewClient(ctx, Config{
		Addrs:    []string{"localhost:6379"},
		Password: "root", // no password set
		Database: 0,      // use default DB
	})
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func TestModel(t *testing.T) {
	rdb, err := NewClient(ctx, Config{
		Addrs:    []string{"localhost:6379"},
		Password: "root", // no password set
		Database: 0,      // use default DB
	})
	if err != nil {
		panic(err)
	}
	m := NewModel(rdb, "test")
	if err = m.Set(ctx, "test", "model", 0).Err(); err != nil {
		panic(err)
	}
	res := m.Get(ctx, "test")
	if res.Err() != nil {
		panic(err)
	}
	fmt.Println(res.String())
}
