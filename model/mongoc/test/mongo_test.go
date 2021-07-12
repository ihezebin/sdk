package test

//
//import (
//	"fmt"
//	"github.com/whereabouts/sdk-go/logger"
//	"github.com/whereabouts/sdk-go/model/mongoc"
//	"testing"
//	"time"
//)
//
//func TestMongo(t *testing.T) {
//	client, err := mongoc.Init(mongoc.Config{
//		Addrs:          []string{"127.0.0.1:27017"},
//		Database:       "test",
//		Username:       "root",
//		Password:       "root",
//		Source:         "admin",
//		ReplicaSetName: "",
//		Timeout:        3,
//		InsertTimeAuto: true,
//		UpdateTimeAuto: true,
//	})
//	defer client.Close()
//	if err != nil {
//		logger.Fatalln(err)
//	}
//	user1 := User{Name: "hezebin1", Age: 21}
//	user2 := User{Name: "hezebin2", Age: 22}
//	user3 := User{Name: "hezebin3", Age: 23}
//	user4 := User{Name: "hezebin4", Age: 24}
//	err = GetUserDB().Add(user1, user2, user3, user4, nil)
//	if err != nil {
//		logger.Errorln(err)
//		return
//	}
//	err = GetUserDB().DeleteTest(24)
//	if err != nil {
//		logger.Errorln(err)
//		return
//	}
//	time.Sleep(3 * time.Second)
//	err = GetUserDB().ModifyAgeByName("hezebin1", 221)
//	if err != nil {
//		logger.Errorln(err)
//		return
//	}
//	u, err0 := GetUserDB().GetByName("hezebin2")
//	if err0 != nil {
//		logger.Errorln(err0)
//		return
//	}
//	fmt.Printf("%+v\n", u)
//	err = GetUserDB().ReplaceByName("hezebin2", user4)
//	if err != nil {
//		logger.Errorln(err)
//		return
//	}
//	users, err1 := GetUserDB().GetAll()
//	if err1 != nil {
//		logger.Errorln(err1)
//		return
//	}
//	for _, user := range users {
//		fmt.Printf("%+v\n", *user)
//	}
//}
