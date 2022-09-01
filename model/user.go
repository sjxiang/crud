package model

import (
	"gorm.io/gorm"
)

// User 用户模型

type User struct {
	gorm.Model
	NickName string `gorm:"type:varchar(20); not null" json:"nickname" binding:"required"`
	Password string `gorm:"type:varchar(20); not null" json:"password" binding:"required"`
	Email    string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
	Phone    string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
}

