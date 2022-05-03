package mongoc

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"testing"
	"time"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	UpdateAt int    `json:"update_at" bson:"update_at"`
	CreateAt int    `json:"create_at" bson:"create_at"`
	DeleteAt int    `json:"delete_at" bson:"delete_at"`
}

func TestRegex(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(ctx, &Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	m := NewBaseModel(c, "test", "user")

	users := make([]*User, 0)
	err = m.FindMany(ctx, bson.M{"name": primitive.Regex{Pattern: regexp.QuoteMeta("A"), Options: "i"}}, &users)
	if err != nil {
		logger.Fatal(err.Error())
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

func TestMongoc(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(ctx, &Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	m := NewBaseModel(c, "test", "user")
	_, err = m.InsertOne(ctx, User{Name: "test", Age: 18})
	if err != nil {
		logger.Fatal(err.Error())
	}
	filter := bson.M{"name": "test"}
	user := User{}
	if err = m.FindOne(ctx, filter, &user); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Infof("%+v", user)
	_, err = m.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"age": 666}})
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err = m.FindOne(ctx, filter, &user); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Infof("%+v", user)
	if _, err = m.DeleteMany(ctx, filter); err != nil {
		logger.Fatal(err.Error())
	}
}

// 610ececb4fb2d58d008e7c58
func TestUpdateNotExist(t *testing.T) {
	ctx := context.TODO()
	c, err := NewClient(ctx, &Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	m := NewBaseModel(c, "test", "user")
	res, err := m.UpdateOne(ctx, bson.M{"id": "123"}, bson.M{"$set": bson.M{"gender": "male"}})
	if err != nil {
		logger.Fatal(err.Error())
	}
	fmt.Println(res)
}

func TestInsert(t *testing.T) {
	c, err := NewClient(context.Background(), &Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	result, err := NewBaseModel(c, "test", "user").InsertOne(context.TODO(), map[string]interface{}{
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
	c, err := NewClient(ctx, &Config{
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
	m := NewBaseModel(c, "test", "user")
	id := "610e657d6c5f1f389720218c"
	user := User{}
	err = m.FindOne(ctx, bson.M{"id": id}, &user)
	if err != nil {
		logger.Errorf("find user err: %s", err.Error())
		return
	}
	logger.Infof("before transaction: %+v", user)

	// failed
	_, err = m.DoWithTransaction(ctx, func(ctx context.Context, model Model) (interface{}, error) {
		updateResult, err := model.UpdateOne(ctx, bson.M{"id": id}, bson.M{"name": "Hezebin1"})
		if err != nil {
			return updateResult, nil
		}
		return nil, errors.New("If the returned err is not nil, The transaction will be rolled back")
	})
	err = m.FindOne(ctx, bson.M{"id": id}, &user)
	if err != nil {
		logger.Errorf("find user err: %s", err.Error())
		return
	}
	logger.Infof("after failed transaction: %+v", user)

	// successful
	_, err = m.DoWithTransaction(ctx, func(ctx context.Context, model Model) (interface{}, error) {
		return model.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"name": "Hezebin2"}})
	})
	if err != nil {
		logger.Errorf("update user err: %s", err.Error())
	}
	err = m.FindOne(ctx, bson.M{"id": id}, &user)
	if err != nil {
		logger.Errorf("find user err: %s", err.Error())
		return
	}
	logger.Infof("after success transaction: %+v", user)
}

func TestAutoTime(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(ctx, &Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	m := NewAutoTimeModel(c, "test", "user")
	_, err = m.InsertOne(ctx, User{Name: "auto_time", Age: 18})
	if err != nil {
		logger.Fatal(err.Error())
	}
	filter := bson.M{"name": "auto_time"}
	user := User{}
	if err = m.FindOne(ctx, filter, &user); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Infof("%+v", user)
	time.Sleep(time.Second * 3)
	_, err = m.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"age": 666}})
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err = m.FindOne(ctx, filter, &user); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Infof("%+v", user)
	if _, err = m.DeleteMany(ctx, filter); err != nil {
		logger.Fatal(err.Error())
	}
}

func TestSoftDelete(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(ctx, &Config{
		Addrs: []string{"127.0.0.1:27017"},
		Auth: &Auth{
			Username: "root",
			Password: "root",
		},
	})
	if err != nil {
		logger.Fatal(err.Error())
	}
	m := NewAutoTimeModel(c, "test", "user").SetSoftDelete(true)
	_, err = m.InsertOne(ctx, User{Name: "soft_delete", Age: 18})
	if err != nil {
		logger.Fatal(err.Error())
	}
	filter := bson.M{"name": "soft_delete"}
	user := User{}
	if err = m.FindOne(ctx, filter, &user); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Infof("%+v", user)
	time.Sleep(time.Second * 3)
	_, err = m.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"age": 666}})
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err = m.FindOne(ctx, filter, &user); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Infof("%+v", user)
	if _, err = m.DeleteMany(ctx, filter); err != nil {
		logger.Fatal(err.Error())
	}
}
