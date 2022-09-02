package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/sjxiang/crud/app/models"
	"github.com/sjxiang/crud/app/requests"
)

// 注册
func (uc UserController) Setup(ctx *gin.Context) {
	var request requests.UserSetupRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "请求解析错误，请确认格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"Error": err.Error(),
		})

		return
	}
	

	// 加密（第二个参数是 cost 值，建议大于 12，数值越大耗费时间越长）
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "加密失败",
			"Error": err.Error(),
		})

		return
	}

	user := models.User{
		NickName: request.NickName,
		Password: string(hash),
		Email: request.Email,
		Phone: request.Phone,
	}
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
