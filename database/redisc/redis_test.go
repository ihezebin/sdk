package redisc

import (
	"fmt"
	"github.com/whereabouts/sdk/logger"
	"testing"
)

func TestRedis(t *testing.T) {
	c, err := NewClientWithConfig(Config{
		Addr:      ":6379",
		Password:  "root",
		MaxIdle:   10,
		MaxActive: 50,
	})
	if err != nil {
		logger.Fatal(err)
	}
	defer c.Close()
	model := NewBaseModel(c, "email")
	res := model.Set("123@qq.com", "123321")
	res = model.Set("123@qq.com", 123)
	fmt.Println("res: ", res)
	res = model.Get("123@qq.com")
	fmt.Println("res: ", res, res)
	res = model.Get("1231231@qq.com")
	fmt.Println("res: ", res)
}
