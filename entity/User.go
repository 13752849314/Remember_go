package entity

import (
	"github.com/goccy/go-json"
	"regexp"
	"remember/common"
)

type User struct {
	ID        uint               `gorm:"primary_key;column:id" json:"id"`
	CreatedAt common.MyDateTime  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt common.MyDateTime  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *common.MyDateTime `sql:"index" gorm:"column:deleted_at" json:"deleted_at"`
	Username  string             `gorm:"column:username;unique" form:"username" json:"username"`
	Password  string             `gorm:"column:password" form:"password" json:"password"`
	Phone     string             `gorm:"column:phone" form:"phone" json:"phone"`
	Email     string             `gorm:"column:email" form:"email" json:"email"`
	Birthday  common.MyDate      `gorm:"column:birthday" form:"birthday" json:"birthday"`
	OpenId    string             `gorm:"column:openId" form:"openId" json:"openId"`
	Roles     string             `gorm:"column:roles" form:"roles" json:"roles"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Desensitization() *User {
	bytes, _ := json.Marshal(u)
	user := new(User)
	_ = json.Unmarshal(bytes, user)
	user.ID = 0
	user.Password = ""
	user.OpenId = ""
	re, _ := regexp.Compile("(\\d{3})\\d{4}(\\d{4})")
	user.Phone = re.ReplaceAllString(user.Phone, "$1****$2")
	return user
}
