package conf


import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/sjxiang/crud/app/models"
)


// 配置初始化
func Init() {

	// 从本地读取环境变量 env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("错误加载 .env：%s\n", err.Error())
	}

	// 连接 MySQL 数据库
	models.SetupDB(os.Getenv("MYSQL_DSN"))

    // 连接 Redis 

}