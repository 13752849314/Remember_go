package impl

import (
	"log"
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
