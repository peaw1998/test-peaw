package calculate

import (
	controller "calculate/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	main := r.Group("/calculate")
	{
		main.GET("/area", controller.Test())
	}
}
