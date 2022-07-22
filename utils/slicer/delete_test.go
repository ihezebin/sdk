package slicer

import (
	"fmt"
	"testing"
)

type Stu struct {
	Name string `json:"name" bson:"_name"`
	Age  int    `json:"age" bson:"_age"`
}

func TestDeleteCone(t *testing.T) {
	s := []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 8, 0}
	fmt.Println("default s:", s)
	s = Delete(s, 8).([]int)
	fmt.Println("s delete one 8:", s)
	stus := []Stu{{"a", 1}, {"b", 2}, {"c", 3}}
	fmt.Println("struct slicer stus:", stus)
	stus = Delete(stus, Stu{"a", 1}).([]Stu)
	fmt.Println("stus delete one Stu{a 1}:", stus)
}

func TestDeleteAll(t *testing.T) {
	s := []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 8, 0}
	fmt.Println("default s:", s)
	s = DeleteAll(s, 1).([]int)
	fmt.Println("s delete all 1:", s)
}
