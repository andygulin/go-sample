package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {
	student := Student{}
	t := reflect.TypeOf(student)
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Name)
		}
	}

	v := reflect.ValueOf(&student)
	v.Elem().FieldByName("Id").SetInt(1)
	v.Elem().FieldByName("Name").SetString("aaaa")
	v.Elem().FieldByName("Age").SetInt(11)

	fmt.Println(student.Id, student.Name, student.Age)
}
