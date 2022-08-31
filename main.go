package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func init() {

	// 从本地读取环境变量 env
	godotenv.Load()
}


type User struct {
	NickName string
	Password string
	Email    string
	Phone    string
}

func main() {

	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}


	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	fmt.Println(err)

	// 连接池配置
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)


	db.AutoMigrate(&User{})


	r := gin.Default()
	r.Run(":"+ os.Getenv("WEB_PORT"))
	
}