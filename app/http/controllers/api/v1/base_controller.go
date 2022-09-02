package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {}


func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Msg": "Pong",
	})
}