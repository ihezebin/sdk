package mongoc

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/whereabouts/sdk-go/logger"
	"testing"
)

type User struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
	Age  int           `json:"age"`
}

func TestMongo(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(Config{
		Addrs:          []string{"127.0.0.1:27017"},
		Database:       "test",
		Username:       "root",
		Password:       "root",
		Source:         "admin",
		ReplicaSetName: "",
		Timeout:        3,
		InsertTimeAuto: true,
		UpdateTimeAuto: true,
	})
	if err != nil {
		logger.Fatalln("create mongodb err:", err)
	}
	defer c.Close()
	users := []User{
		{Name: "korbin", Age: 18},
		{Name: "hezebin", Age: 19},
		{Name: "whereabouts", Age: 20},
		{Name: "nilName", Age: 21},
	}
	db := NewMongoDB(c, "test", "user")
	err = db.Insert(ctx, users)
	if err != nil {
		logger.Errorln("insert user err: ", err)
		return
	}
	err = db.Remove(ctx, bson.M{"name": "nilName"})
	if err != nil {
		logger.Errorln("delete user which name is nilName err:", err)
		return
	}
	user := User{}
	err = db.Modify(ctx, bson.M{"name": "korbin"}, bson.M{"age": 180}, &user)
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
