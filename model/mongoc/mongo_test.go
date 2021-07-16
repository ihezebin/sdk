package mongoc

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/whereabouts/sdk-go/logger"
	"testing"
	"time"
)

type User struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name"`
	Age        int           `json:"age"`
	CreateTime string        `json:"create_time"`
	UpdateTime string        `json:"update_time"`
}

func TestMongo(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(Config{
		Addrs:          []string{"127.0.0.1:27017"},
		Username:       "root",
		Password:       "root",
		Source:         "admin",
		ReplicaSetName: "",
		Timeout:        3,
		AutoTime:       true,
	})
	if err != nil {
		logger.Fatalln("create mongodb err:", err)
	}
	defer c.Close()
	db := NewMongoDB(c, "test", "user")
	users := []interface{}{
		&User{Name: "korbin", Age: 18},
		&User{Name: "hezebin", Age: 19},
		&User{Name: "whereabouts", Age: 20},
		&User{Name: "nilName", Age: 21},
	}
	err = db.Insert(ctx, users...)
	if err != nil {
		logger.Errorln("insert user err: ", err)
		return
	}
	err = db.RemoveOne(ctx, bson.M{"name": "nilName"})
	if err != nil {
		logger.Errorln("delete user which name is nilName err:", err)
		return
	}
	time.Sleep(time.Second * 2)
	err = db.ModifyOne(ctx, bson.M{"name": "korbin"}, bson.M{"age": 180})
	if err != nil {
		logger.Errorln("update user korbin age 18 to 180 err:", err)
		return
	}
	userList := make([]*User, 0)
	err = db.FindAll(ctx, nil, nil, nil, 0, 0, &userList)
	if err != nil {
		logger.Errorln("get user list err:", err)
		return
	}
	for _, u := range userList {
		fmt.Println(u)
	}
}

func TestModifyOne(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(Config{
		Addrs:          []string{"127.0.0.1:27017"},
		Username:       "root",
		Password:       "root",
		Source:         "admin",
		ReplicaSetName: "",
		Timeout:        3,
		AutoTime:       true,
	})
	if err != nil {
		logger.Fatalln("create mongodb err:", err)
	}
	defer c.Close()
	db := NewMongoDB(c, "test", "user")
	err = db.ModifyOne(ctx, nil, bson.M{"age": 777})
	if err != nil {
		logger.Errorln("modify one err:", err)
		return
	}
}

func TestReplaceOne(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(Config{
		Addrs:          []string{"127.0.0.1:27017"},
		Username:       "root",
		Password:       "root",
		Source:         "admin",
		ReplicaSetName: "",
		Timeout:        3,
		AutoTime:       true,
	})
	if err != nil {
		logger.Fatalln("create mongodb err:", err)
	}
	defer c.Close()
	db := NewMongoDB(c, "test", "user")
	err = db.ReplaceOne(ctx, nil, nil)
	if err != nil {
		logger.Errorln("modify one err:", err)
		return
	}
}

func TestFindOne(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(Config{
		Addrs:          []string{"127.0.0.1:27017"},
		Username:       "root",
		Password:       "root",
		Source:         "admin",
		ReplicaSetName: "",
		Timeout:        3,
		AutoTime:       true,
	})
	if err != nil {
		logger.Fatalln("create mongodb err:", err)
	}
	defer c.Close()
	db := NewMongoDB(c, "test", "user")
	user := &User{}
	//err = db.FindOne(ctx, nil, []string{}, user)
	err = db.FindOne(ctx, nil, nil, user)
	if err != nil {
		logger.Errorln("find one err:", err)
		return
	}
	fmt.Println(user)
}
