package main

import (
	"fmt"
	"reflect"
	"remember/config"
	"remember/entity"
	"remember/mapper"
)

func main() {
	fmt.Println("Hello Remember_go!")

	fmt.Println(config.Configure)

	xml := mapper.ReadXml()
	fmt.Println(xml)

	um := mapper.GetMapper("userMapper")
	fmt.Println(um)

	a := reflect.TypeOf(entity.User{})
	b := reflect.New(a)
	fmt.Println(a, b)

	mp := mapper.UserMapper{}
	users := mp.Select("getAllUsers")
	ok := users.(*[]entity.User)
	fmt.Println(ok)
	fmt.Println(users)
}
