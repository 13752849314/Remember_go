package entity

import (
	"remember/common"
)

type Bill struct {
	ID           uint               `gorm:"primary_key;column:id" json:"id"`
	CreatedAt    common.MyDateTime  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt    common.MyDateTime  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt    *common.MyDateTime `sql:"index" gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	Username     string             `gorm:"column:username" json:"username"`
	ConsumeType  int                `gorm:"column:consumeType" json:"consumeType"`
	ConsumeMoney float64            `gorm:"column:consumeMoney" json:"consumeMoney"`
	ConsumeTime  common.MyDateTime  `gorm:"column:consumeTime;type:datetime" json:"consumeTime"`
	Remark       string             `gorm:"column:remark" json:"remark"`
}

func (b *Bill) TableName() string {
	return "bills"
}
