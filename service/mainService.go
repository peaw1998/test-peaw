package service

func Resp(c *gin.Context, status int, msg string, data interface{}) {
	var code int = 1
	var statusCode int = 200
	var statusMsg string = "ok"
	// if status == 0 || status == 200 {
	// 	code = 0
	// } else if status == 101 {
	// 	code = 101
	// } else if status == 1 {
	// 	code = 1
	// } else if status == 400 {
	// 	statusCode = 400
	// 	statusMsg = "invalid data"
	// } else if status == 401 {
	// 	statusCode = 401
	// 	c.AbortWithStatusJSON(statusCode, gin.H{
	// 		"status": map[string]interface{}{
	// 			"code":    statusCode,
	// 			"message": "unauthorized",
	// 		},
	// 		"message": map[string]interface{}{
	// 			"code":    1,
	// 			"message": "unauthorized",
	// 		},
	// 		"data": data,
	// 	})
	// 	return
	// } else if status == 500 {
	// 	statusCode = 500
	// 	statusMsg = "server error"
	// }

	c.JSON(statusCode, gin.H{
		"status": map[string]interface{}{
			"code":    statusCode,
			"message": statusMsg,
		},
		"message": map[string]interface{}{
			"code":    code,
			"message": msgGlobal,
		},
		"data": data,
	})
	return
}