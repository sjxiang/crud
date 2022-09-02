package models

import (
	"gorm.io/gorm"
)

// User 用户模型

type User struct {
	gorm.Model
	NickName string `gorm:"type:varchar(20);column:nick_name;not null"`
	Password string `gorm:"type:varchar(60);not null"`
	Email    string `gorm:"type:varchar(40);unique"`
	Phone    string `gorm:"type:varchar(20);unique"`
}


// TODO 校验参数标签 "github.com/thedevsaddam/govalidator"