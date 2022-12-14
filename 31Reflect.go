package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id     string
	Parent []string
	Name   string
	Age    int
}

type UserBody struct {
	Name string
	Age  int
}

func Reflect() {
	u1 := User{"1", []string{"yakuza", "suzu"}, "yuza", 12}
	fmt.Println(u1)
	fmt.Println("-------------")

	fields := reflect.VisibleFields(reflect.TypeOf(u1))
	r := reflect.ValueOf(u1).FieldByName("Name")
	if r.IsValid() {
		fmt.Println(r)
	} else {
		fmt.Println("r is invalid")
    }

	fmt.Println("-------------")

	fmt.Println(r)
	for i, field := range fields {
		fmt.Println(i, field.Name)
	}
}
