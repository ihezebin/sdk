package v1

import (
	"fmt"
	"github.com/ihezebin/sdk/logger"
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
	res := model.Set("123@qq.com", "123")
	fmt.Println("res: ", res.Error())
	fmt.Println("res: ", model.SetIfNotExists("123@qq.com", 123))
	res = model.Get("123@qq.com")
	fmt.Println("res: ", res.String())
	res = model.Get("1231231@qq.com")
	fmt.Println("res: ", res.String())
	res = model.IncrBy("123@qq.com", 2)
	fmt.Println("res: ", res.Int())
}
