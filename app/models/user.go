package models

import (
	"gorm.io/gorm"
)

// User 用户模型

type User struct {
	gorm.Model
	NickName string `gorm:"type:varchar(20);column:nick_name;not null" json:"nickname"`
	Password string `gorm:"type:varchar(60);not null"                  json:"password"`  // TODO 敏感信息，json 转义忽略
	Email    string `gorm:"type:varchar(40);unique"                    json:"email"`
	Phone    string `gorm:"type:varchar(20);unique"                    json:"phone"`
}


// TODO 校验参数标签 "github.com/thedevsaddam/govalidator"