package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/crud/app/models"
)

func (uc UserController) Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	nickname := user.(models.User).NickName

	ctx.JSON(http.StatusOK, gin.H{
		"Msg": fmt.Sprintf("%v 在线", nickname),
	})
}