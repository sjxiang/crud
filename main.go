package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/crud/conf"
	"github.com/sjxiang/crud/controllers/api/v1/user"
)


func main() {

	// 从配置文件读取配置
	conf.Init()

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