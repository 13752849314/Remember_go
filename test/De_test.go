package test

import (
	"fmt"
	"remember/utils"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	Str := "admin"
	pass := utils.PasswordEncrypt(Str)
	fmt.Println(Str)
	fmt.Println(pass)
}
