package mapper

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"remember/database"
	"remember/entity"
)

var dbBill *gorm.DB
var umBill map[string]Select

type BillMapper struct {
}

func init() {
	dbBill = database.GetCoon()
	umBill = GetMapper("billMapper")
}

func (b *BillMapper) Insert(bill *entity.Bill) error {
	err := dbBill.Create(bill).Error
	return err
}

func (b *BillMapper) Update(bill *entity.Bill) error {
	err := dbBill.Save(bill).Error
	return err
}

func (b *BillMapper) Delete(bill *entity.Bill) error {
	err := dbBill.Delete(bill).Error
	return err
}

func (b *BillMapper) Select(funName string, where ...interface{}) interface{} {
	s := umBill[funName]
	Sql := s.Sql
	reType := s.ResultType
	Type := GetType(reType)
	ret := reflect.New(Type).Interface()
	dbUser.Raw(Sql, where...).Scan(ret)
	return ret
}
