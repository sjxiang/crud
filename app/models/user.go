package models

import (
	"gorm.io/gorm"
)

// User 用户模型

type User struct {
	gorm.Model

	UserName       string `gorm:"type:varchar(20);not null"` 
	PasswordDigest string `gorm:"type:varchar(60);not null"`  // 密码摘要
	Email          string `gorm:"type:varchar(40);unique"`
	Phone          string `gorm:"type:varchar(20);unique"`
	Status         string `gorm:"type:varchar(20)"`
	Avatar         string `gorm:"size:1000"` // 头像
}


const (
	// Active 激活用户
	Active string = "active"
	
	// Inactive 未激活用户
	Inactive string = "inactive"
	
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// TODO 校验参数标签 "github.com/thedevsaddam/govalidator"