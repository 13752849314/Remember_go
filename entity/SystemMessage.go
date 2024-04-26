package entity

import "remember/common"

type SystemMessage struct {
	ID          uint               `gorm:"primary_key;column:id" json:"id"`
	CreatedAt   common.MyDateTime  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   common.MyDateTime  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt   *common.MyDateTime `sql:"index" gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	PublishTime common.MyDateTime  `gorm:"publish_time;type:datetime" json:"publish_time"`
	Publisher   string             `gorm:"publisher" json:"publisher"`
	Message     string             `gorm:"message" json:"message"`
}

func (s *SystemMessage) TableName() string {
	return "systemMessages"
}
