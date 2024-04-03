package test

import (
	"fmt"
	"remember/entity"
	"remember/utils"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	Str := "admin"
	pass := utils.PasswordEncrypt(Str)
	fmt.Println(Str)
	fmt.Println(pass)
}

func TestCreatToken(t *testing.T) {
	user := new(entity.User)
	user.Username = "user"
	user.ID = 1
	token, err := utils.CreatToken(user)
	if err != nil {
		return
	}
	fmt.Println("token", token)
	checkToken, err := utils.CheckToken(token)
	if err != nil {
		return
	}
	fmt.Println(checkToken)
}

func TestMap(t *testing.T) {
	login := make(map[string]string)
	login["123"] = "123"
	s := login["234"]
	fmt.Println(s)
}
