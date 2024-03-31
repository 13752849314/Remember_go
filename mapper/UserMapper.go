package mapper

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"remember/database"
	"remember/entity"
)

var DBUser *gorm.DB
var UM map[string]Select

type UserMapper struct {
}

func init() {
	DBUser = database.GetCoon()
	UM = GetMapper("userMapper")
}

func (m *UserMapper) Insert(user *entity.User) error {
	err := DBUser.Create(user).Error
	return err
}

func (m *UserMapper) Update(user *entity.User) error {
	err := DBUser.Update(user).Error
	return err
}

func (m *UserMapper) Delete(user *entity.User) error {
	err := DBUser.Delete(user).Error
	return err
}

func (m *UserMapper) GetUserById(id int) *entity.User {
	user := new(entity.User)
	DBUser.Find(user, id)
	return user
}

func (m *UserMapper) Select(funName string, where ...interface{}) interface{} {
	s := UM[funName]
	Sql := s.Sql
	reType := s.ResultType
	Type := GetType(reType)
	ret := reflect.New(Type).Interface()
	DBUser.Raw(Sql, where...).Scan(ret)
	return ret
}
