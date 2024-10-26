package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello world!",
	})
}
