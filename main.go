package main

import (
	"os"

	"github.com/gin-gonic/gin"

	base "github.com/sjxiang/crud/app/http/controllers/api/v1"
	"github.com/sjxiang/crud/app/http/controllers/api/v1/user"
	"github.com/sjxiang/crud/app/http/middlewares"
	"github.com/sjxiang/crud/conf"
)



func main() {

	// 从配置文件读取配置
	conf.Init()

	r := gin.Default()

	// 测试
	r.GET("/ping", base.Ping)


	v1 := r.Group("/api/v1")
	{
		uc := new(user.UserController)

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

	
		// 用户注册
		v1.POST("/user/signup", uc.Setup)

		// 用户登录
		v1.POST("/user/login", uc.Login)

		// 身份验证
		v1.GET("/user/validate", middlewares.RequireAuth, uc.Validate)

	}

	
	r.Run(":"+ os.Getenv("WEB_PORT"))
	
}