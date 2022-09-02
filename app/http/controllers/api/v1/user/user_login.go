package user

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/sjxiang/crud/app/models"
	"github.com/sjxiang/crud/app/requests"
)

// 登录
func (uc UserController) Login(ctx *gin.Context) {

	// 获得请求正文（邮箱、密码）
	var request requests.UserLoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "请求解析错误，请确认格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"Error": err.Error(),
		})

		return
	}

	// 查询请求登录的用户
	var user models.User
	models.DB.Where("email = ?", request.Email).Find(&user)
	// models.DB.First(&user, "email = ?", request.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "无效的 email",
		})

		return
	}

	// 将请求正文里的密码与保存的用户密码哈希值进行比较
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "无效的 password",
		})

		return
	}

	// 生成 jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// 签名（私钥）
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Msg": "创建 token 失败",
			"Error": err.Error(),
		})

		return
	}

	// 返回
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{})
}
