package model

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
	"gorm.io/gorm"
)

var DB *gorm.DB


// 初始化 MySQL 连接
func SetupDB(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 解决表名复数，user
		  },
	})

	if err != nil {
		log.Panicf("连接 db 失败：%s\n", err.Error())
	}


	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}

	// 连接池配置

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	DB = db

	migration()
}