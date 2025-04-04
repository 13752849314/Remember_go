package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"remember/entity"
)

var login map[string]string

func init() {
	login = make(map[string]string)
}

func AddInfo(username, token string) error {
	s := login[username]
	if s != "" {
		login[username] = s
		return errors.New("用户已经登录，请勿重复登录\n" + "token:" + s)
	}
	login[username] = token
	log.Println("用户：", username, "成功登录")
	return nil
}

func GetInfo(username string) string {
	return login[username]
}

func DeleteInfo(username, token string) error {
	t := GetInfo(username)
	if t == token {
		delete(login, username)
		return nil
	} else {
		return errors.New("用户已经退出")
	}
}

func DeleteInfoByUsername(username string) {
	delete(login, username)
}

func GetKeys() []string {
	keys := make([]string, 0)
	for key := range login {
		keys = append(keys, key)
	}
	return keys
}

func GetUser(c *gin.Context) *entity.User {
	user, _ := c.Get("user")
	return user.(*entity.User)
}

func MaintainLogin() {
	if len(login) == 0 {
		return
	}
	deleted := make([]string, 0)
	for k, v := range login {
		_, err := CheckToken(v)
		if err != nil {
			deleted = append(deleted, k)
		}
	}
	if len(deleted) == 0 {
		return
	}
	for _, s := range deleted {
		delete(login, s)
	}
	log.Println("当前在线用户：", GetKeys())
}
