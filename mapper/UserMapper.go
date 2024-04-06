package mapper

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"remember/database"
	"remember/entity"
)

var dbUser *gorm.DB
var umUser map[string]Select

type UserMapper struct {
}

func init() {
	dbUser = database.GetCoon()
	umUser = GetMapper("userMapper")
}

func (m *UserMapper) Insert(user *entity.User) error {
	err := dbUser.Create(user).Error
	return err
}

func (m *UserMapper) Update(user *entity.User) error {
	err := dbUser.Save(user).Error
	return err
}

func (m *UserMapper) Delete(user *entity.User) error {
	err := dbUser.Delete(user).Error
	return err
}

func (m *UserMapper) GetAllUser() []entity.User {
	var users []entity.User
	dbUser.Find(&users)
	return users
}

func (m *UserMapper) GetUserById(id int) *entity.User {
	user := new(entity.User)
	dbUser.Find(user, id)
	return user
}

func (m *UserMapper) GetUserByUsername(username string) *entity.User {
	user := new(entity.User)
	dbUser.First(user, "username= ? ", username)
	return user
}

func (m *UserMapper) Select(funName string, where ...interface{}) interface{} {
	s := umUser[funName]
	Sql := s.Sql
	reType := s.ResultType
	Type := GetType(reType)
	ret := reflect.New(Type).Interface()
	dbUser.Raw(Sql, where...).Scan(ret)
	return ret
}
