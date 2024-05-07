package mapper

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"remember/database"
	"remember/entity"
)

var dbSystemMessage *gorm.DB
var umSystemMessage map[string]Select

type SystemMessageMapper struct {
}

func init() {
	dbSystemMessage = database.GetCoon()
	umSystemMessage = GetMapper("systemMessageMapper")
}

func (m *SystemMessageMapper) Insert(message *entity.SystemMessage) error {
	err := dbSystemMessage.Create(message).Error
	return err
}

func (m *SystemMessageMapper) Update(message *entity.SystemMessage) error {
	err := dbSystemMessage.Save(message).Error
	return err
}

func (m *SystemMessageMapper) Delete(message *entity.SystemMessage) error {
	err := dbSystemMessage.Delete(message).Error
	return err
}

func (m *SystemMessageMapper) Select(funName string, where ...interface{}) interface{} {
	s := umSystemMessage[funName]
	Sql := s.Sql
	reType := s.ResultType
	Type := GetType(reType)
	ret := reflect.New(Type).Interface()
	dbUser.Raw(Sql, where...).Scan(ret)
	return ret
}
