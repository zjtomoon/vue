package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Name     string                 // 姓名
	Email    string                 // 邮箱
	Mobile   string `gorm:"unique"` // 手机
	QQ       string
	Gender   int // 性别 0男1女
	Age      int // 年龄
	Remark   string // 备注
	Token    string `gorm:"-"`
	Session  string `gorm:"-"`
}
