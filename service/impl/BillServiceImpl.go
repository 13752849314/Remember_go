package impl

import (
	"errors"
	"log"
	"reflect"
	"remember/entity"
	"remember/mapper"
)

type BillServiceImpl struct {
	mapper mapper.BillMapper
}

func (b *BillServiceImpl) GetBillsByUsername(username string) interface{} {
	i := b.mapper.Select("getBillsByUsername", username)
	return i
}

func (b *BillServiceImpl) AddBill(bill *entity.Bill) error {
	err := b.mapper.Insert(bill)
	if err == nil {
		log.Println("账单添加：", bill)
	}
	return err
}

func (b *BillServiceImpl) DeleteBillById(id int) error {
	bill, err := b.mapper.GetBillById(id)
	if err != nil {
		return errors.New("账单不存在")
	}
	err = b.mapper.Delete(bill)
	return err
}

func (b *BillServiceImpl) ChangeBillInfoById(id int, mp map[string]any) error {
	bill, err := b.mapper.GetBillById(id)
	if bill.ID == 0 {
		return errors.New("账单不存在")
	}
	if err != nil {
		return err
	}
	v := reflect.ValueOf(bill)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for key, value := range mp {
		if reflect.ValueOf(value).Type() == v.FieldByName(key).Type() {
			v.FieldByName(key).Set(reflect.ValueOf(value))
		}
	}
	err = b.mapper.Update(bill)
	if err != nil {
		return err
	}
	return nil
}
