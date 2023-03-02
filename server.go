package calculate

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.Use(CORS)

	Route(r)

	r.Run()
}


func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
		return
	} else {
		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		// c.AbortWithStatusJSON(401, gin.H{
		// 	"code":    "ERROR",
		// 	"message": "Unauthorized",
		// })
		c.AbortWithStatus(http.StatusOK)
		return
	}
}