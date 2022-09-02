package user

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/crud/app/models"
)

// 注册
func (uc UserController) Setup(ctx *gin.Context) {
	
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "请求解析错误，请确认格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"Error": err.Error(),
		})

		return
	}
	

	// 加密（第二个参数是 cost 值，建议大于 12，数值越大耗费时间越长）
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "加密失败",
			"Error": err.Error(),
		})

		return
	}

	user.Password = string(hash)
	result := models.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "创建用户失败",
			"Error": result.Error,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Msg": "注册成功",
		"Data": user,
	})






}



// var body struct {
// 	Email    string
// 	Password string
// }

// if c.Bind(&body) != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"Error": "读取 body 失败",
// 	})

// 	return
// }