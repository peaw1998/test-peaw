package calculate

import (
	"github.com/gin-gonic/gin"
)

// MemberRoute create route
func Route(r *gin.Engine, resource *db.Resource) {
	main := r.Group("/test")
	{
		main.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		} )
	}
}
