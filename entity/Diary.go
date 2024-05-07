package entity

import "remember/common"

type Diary struct {
	ID         uint               `gorm:"primary_key;column:id" json:"id"`
	CreatedAt  common.MyDateTime  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt  common.MyDateTime  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt  *common.MyDateTime `sql:"index" gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	RecordDate common.MyDate      `gorm:"column:recordDate;type:datetime" json:"recordDate"`
	Mood       string             `gorm:"column:mood" json:"mood"`
	Username   string             `gorm:"column:username" json:"username"`
	Weather    string             `gorm:"column:weather" json:"weather"`
	Type       string             `gorm:"column:type" json:"type"`
	Message    string             `gorm:"column:message" json:"message"`
}

func (d *Diary) TableName() string {
	return "diary"
}
