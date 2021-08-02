package mongo

import (
	"context"
	"github.com/whereabouts/sdk/logger"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestInsert(t *testing.T) {
	c, err := NewClient(context.Background(), Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	id, err := c.NewBaseModel("test", "user").Insert(context.TODO(), User{
		Name: "aaa",
		Age:  111,
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(id)
}
