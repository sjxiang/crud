package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sjxiang/crud/model"
	"github.com/sjxiang/crud/controllers/api/v1/user"

)


func init() {

	// 从本地读取环境变量 env
	godotenv.Load()

	// 初始化 MySQL 数据库连接
	dsn := os.Getenv("MYSQL_DSN")
	model.SetupDB(dsn)
}


func main() {

	r := gin.Default()

	// 测试
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Msg": "Pong",
		})
	})


	v1 := r.Group("/api/v1")
	{
		uc := new(user.VideoController)

		// 增
		v1.POST("/user/add", uc.CreateUser)
		// 删
		v1.DELETE("/user/delete/:id", uc.DeleteUser)
		// 改
		v1.PUT("/user/update/:id", uc.UpdateUser)
		// 查（条件）
		v1.GET("/user/list/:nickname", uc.ShowUser)
		// 查（分页、批量）
		v1.GET("/user/list", uc.BatchShowUser)
	}

	
	r.Run(":"+ os.Getenv("WEB_PORT"))
	
}