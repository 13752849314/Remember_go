package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Bill struct {
	gorm.Model
	Username     string    `gorm:"column:username" json:"username"`
	ConsumeType  int       `gorm:"column:consumeType" json:"consumeType"`
	ConsumeMoney float64   `gorm:"column:consumeMoney" json:"consumeMoney"`
	ConsumeTime  time.Time `gorm:"column:consumeTime" json:"consumeTime"`
	Remark       string    `gorm:"column:remark" json:"remark"`
}

func (b *Bill) TableName() string {
	return "bills"
}
