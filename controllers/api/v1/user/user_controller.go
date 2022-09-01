package user

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"


	"github.com/sjxiang/crud/controllers"
	"github.com/sjxiang/crud/models"

)


type UserController struct {
	controllers.BaseController
}


// 增
func (uc UserController) CreateUser(ctx *gin.Context) {
	var data models.User
		
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": 400,  // 400 请求格式错误
			"Msg": "添加失败，提交 JSON 格式错误",
		})

		return
	}

	// 数据库操作 持久化
	models.DB.Create(&data)  // 创建 1 条数据

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "添加成功",
		"Data": data,
	})
}


// 删
func (uc UserController) DeleteUser(ctx *gin.Context) {

	// 1. 找到 id 所对应的记录
	// 2. 判断 id 是否存在
	// 3. 从数据库中删除 / 返回 id 没有找到
	var data []models.User

	id := ctx.Param("id")

	models.DB.Where("id = ?", id).Find(&data)

	if len(data) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{  // 数据库查询失败 或者 没有这个记录
			"Code": 400,
			"Msg": "用户 id 没有找到，删除失败",
		})

		return
	}
	
	models.DB.Where("id = ?", id).Delete(&data)

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200,
		"Msg": "删除成功",
	})
}


// 
func (uc UserController) UpdateUser(ctx *gin.Context) {

	var data models.User

	// 接受 id
	id := ctx.Param("id")

	// 
	models.DB.Select("id").Where("id = ?", id).Find(&data)  // SELECT `id` WHERE `id` = ? FROM `user`;
	
	// 判断 id 是否存在
	if data.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": 400,
			"Msg": "用户 id 没有找到",
		})

		return
	}
	

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": 400,
			"Msg": "修改失败，提交 JSON 格式错误",
		})

		return
	}

	models.DB.Where("id = ?", id).Updates(&data)  // 好几种写法

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200, 
		"Msg": "修改成功",
		"Data": data,
	})
}


// 查（条件）
func (uc UserController) ShowUser(ctx *gin.Context) {

	var data []models.User

	nickname := ctx.Param("nickname")

	models.DB.Where("nick_name = ?", nickname).Find(&data)

	if len(data) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": 400,
			"Msg": "查询失败，查无此人",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg": "查询成功",
		"Data": data,
	})
}


// 查（分页、批量）
func (uc UserController) BatchShowUser(ctx *gin.Context) {

	var data []models.User

	// 查询分页数据 
	// e.g. ?pageSize=10&pageNum=1 第 1 页返回 10 条数据
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum")) 


	// 判断是否需要分页
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageSize = -1
	}

	offsetVal := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}

	// 查询 1 个总数
	var total int64
	models.DB.Model(data).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&data)

	if len(data) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code": 400, 
			"Msg": "查询失败，没有查询到数据",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code": 200, 
		"Msg": "查询成功",
		"Data": gin.H{
			"Users": data, 
			"total": total, 
			"pageNum": pageNum,
			"pageSize": pageSize,
		},
	})
}