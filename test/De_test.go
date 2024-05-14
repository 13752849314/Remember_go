package test

import (
	"fmt"
	"reflect"
	"remember/common"
	"remember/entity"
	"remember/utils"
	"testing"
	"time"
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

func TestStruct2Map(t *testing.T) {
	c := new(common.ChangeUserP)
	c.NewPassword = "123"
	now := time.Now()
	d := common.ChangeUserI{
		Phone:    "111",
		Birthday: common.MyDate(now),
	}

	struct2Map := utils.Struct2Map(c)
	struct2Map1 := utils.Struct2Map(d)

	fmt.Println(struct2Map)
	fmt.Println(struct2Map1)

	v := reflect.ValueOf(&d).Elem()
	name := v.FieldByName("Email")
	name.Set(reflect.ValueOf("1233333"))
	fmt.Println(name)
	time.Sleep(1 * time.Second)
	n := common.MyDate(time.Now())
	v.FieldByName("Birthday").Set(reflect.ValueOf(n))

	fmt.Println(d)
}

func TestFile(t *testing.T) {
	s := utils.ListDir("../entity")
	fmt.Println(s)
	for _, info := range s {
		fmt.Println(info.Name(), info.Size())
	}
}
