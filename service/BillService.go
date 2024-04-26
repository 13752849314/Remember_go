package service

import "remember/entity"

type BillService interface {
	GetBillsByUsername(username string) interface{}
	AddBill(bill *entity.Bill) error
	DeleteBillById(id int) error
	ChangeBillInfoById(id int, mp map[string]any) error
}
