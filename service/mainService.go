package service

import (
	"github.com/gin-gonic/gin"
)

func Resp(c *gin.Context, status int, msg string, data interface{}) {
	var code int = 1
	var statusCode int = 200
	var statusMsg string = "ok"

	c.JSON(statusCode, gin.H{
		"status": map[string]interface{}{
			"code":    statusCode,
			"message": statusMsg,
		},
		"message": map[string]interface{}{
			"code":    code,
			"message": "",
		},
		"data": data,
	})
	return
}
