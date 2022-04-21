package v2

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"testing"
)

type User struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name"`
	Age  int                `json:"age"`
}

func TestRegex(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(ctx, Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	model := c.NewBaseModel("test", "user")

	users := make([]*User, 0)
	err = model.Find(ctx, bson.M{"name": primitive.Regex{Pattern: regexp.QuoteMeta("A"), Options: "i"}}, &users)
	if err != nil {
		logger.Fatal(err.Error())
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

func TestMongoc(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(ctx, Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	base := c.NewBaseModel("test", "user")
	result, err := base.UpdateId(ctx, "60ed5a48e00ed00b91b25000", bson.M{"$set": bson.M{"age": 666}})
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("%+v", result)
}

// 610ececb4fb2d58d008e7c58
func TestUpdateNotExist(t *testing.T) {
	ctx := context.TODO()
	c, err := NewClient(ctx, Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	model := c.NewBaseModel("test", "user")
	res, err := model.UpdateOne(ctx, bson.M{"id": "123"}, bson.M{"$set": bson.M{"gender": "male"}})
	if err != nil {
		logger.Fatal(err.Error())
	}
	fmt.Println(res)
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
	result, err := c.NewBaseModel("test", "user").InsertOne(context.TODO(), map[string]interface{}{
		"_id":  primitive.NewObjectID().Hex(),
		"name": "Korbin",
		"age":  18,
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(result.InsertedID)
}

// ObjectId("610e657d6c5f1f389720218c")
// === RUN   TestTransaction
// {"file":"model_test.go:59","func":"mongoc.TestTransaction","level":"info","msg":"before transaction: {Name:Korbin Age:18}","time":"2021-08-07 18:51:05"}
// {"file":"model_test.go:74","func":"mongoc.TestTransaction","level":"info","msg":"after failed transaction: {Name:Korbin Age:18}","time":"2021-08-07 18:51:05"}
// {"file":"model_test.go:88","func":"mongoc.TestTransaction","level":"info","msg":"after success transaction: {Name:Hezebin2 Age:18}","time":"2021-08-07 18:51:05"}
// --- PASS: TestTransaction (0.04s)
//PASS
func TestTransaction(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(ctx, Config{
		Addrs: []string{"127.0.0.1:27018"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
		ReplicaSetName: "winrs",
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	model := c.NewBaseModel("test", "user")
	id := "610e657d6c5f1f389720218c"
	user := User{}
	err = model.FindId(ctx, id, &user)
	if err != nil {
		logger.Errorf("find user err: %s", err.Error())
		return
	}
	logger.Infof("before transaction: %+v", user)

	// failed
	_, err = model.DoWithTransaction(ctx, func(ctx context.Context, model *Base) (interface{}, error) {
		updateResult, err := model.UpdateId(ctx, id, bson.M{"name": "Hezebin1"})
		if err != nil {
			return updateResult, nil
		}
		return nil, errors.New("If the returned err is not nil, The transaction will be rolled back")
	})
	err = model.FindId(ctx, id, &user)
	if err != nil {
		logger.Errorf("find user err: %s", err.Error())
		return
	}
	logger.Infof("after failed transaction: %+v", user)

	// successful
	_, err = model.DoWithTransaction(ctx, func(ctx context.Context, model *Base) (interface{}, error) {
		return model.UpdateId(ctx, id, bson.M{"$set": bson.M{"name": "Hezebin2"}})
	})
	if err != nil {
		logger.Errorf("update user err: %s", err.Error())
	}
	err = model.FindId(ctx, id, &user)
	if err != nil {
		logger.Errorf("find user err: %s", err.Error())
		return
	}
	logger.Infof("after success transaction: %+v", user)
}
