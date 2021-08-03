package mapper

import (
	"fmt"
	"log"
	"testing"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string
	Weight int `json:"-"`
}

type Stu struct {
	Name string `json:"name" bson:"_name"`
	Age  int    `json:"age" bson:"_age"`
}

func TestStruct2Map(t *testing.T) {
	m, err := Struct2Map(Person{
		Name:   "Korbin",
		Age:    22,
		Gender: "男",
		Weight: 50,
	})
	if err != nil {
		log.Println("struct to map err:", err)
		return
	}
	fmt.Println(m)
	fmt.Println(Json2Map(`{"name": "korbin", "age": 22}`))
	fmt.Println(Map2Json(map[string]string{"name": "korbin", "age": "22"}))
	stu := Stu{}
	err = Map2Struct(map[string]interface{}{"name": "korbin", "age": 22}, &stu)
	if err != nil {
		log.Println("map to struct err:", err)
	}
	fmt.Println(stu)
}
