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

func TestStruct2Map(t *testing.T) {
	m, err := Struct2Map(Person{
		Name:   "Korbin",
		Age:    22,
		Gender: "男",
		Weight: 50,
	})
	if err != nil {
		log.Printf("struct to map err: %v", err)
		return
	}
	fmt.Println(m)
}
