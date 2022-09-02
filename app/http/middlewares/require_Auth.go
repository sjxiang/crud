package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sjxiang/crud/app/models"
)


func RequireAuth(ctx *gin.Context) {

	// 获取 cookie 的请求
	tokenString, err := ctx.Cookie("Authorization")
	
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// 解码 
	// 参数 1. tokenString 
	// 参数 2. 解析私钥的回调函数
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	// 验证
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// 检查 expired time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// 找到带有 token sub 标记的 user
		var user models.User
		models.DB.First(&user, "id = ?", claims["sub"])

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// 附加到 req
		ctx.Set("user", user)

		ctx.Next()

	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)		
	}
}
