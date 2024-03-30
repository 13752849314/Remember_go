package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string     `gorm:"column:username;unique" form:"username" json:"username"`
	Password string     `gorm:"column:password" form:"password" json:"password"`
	Phone    string     `gorm:"column:phone" form:"phone" json:"phone"`
	Email    string     `gorm:"column:email" form:"email" json:"email"`
	Birthday *time.Time `gorm:"column:birthday" form:"birthday" time_format:"2006-01-02" json:"birthday"`
	OpenId   string     `gorm:"column:openId" form:"openId" json:"openId"`
	Roles    string     `gorm:"column:roles" form:"roles" json:"roles"`
}

func (u *User) TableName() string {
	return "users"
}
