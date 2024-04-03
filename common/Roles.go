package common

import (
	"log"
	"reflect"
	"runtime"
	"strings"
)

type Roles struct {
	GetAllUser  string `roles:"admin"`
	GetAllUsers string `roles:"admins"`
}

func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	splits := strings.Split(funcName, ".")
	return splits[len(splits)-1]
}

func GetFuncRoles() string {
	funcName := GetFuncName()
	log.Println("Call:", funcName)
	of := reflect.TypeOf(Roles{})
	name, b := of.FieldByName(funcName)
	if !b {
		return "user"
	}
	role := name.Tag.Get("roles")
	return role
}
