package main

import (
	"fmt"
	"remember/entity"
)

func main() {
	fmt.Println("Hello Remember_go!")

	user := new(entity.User)
	fmt.Println(user)
}
