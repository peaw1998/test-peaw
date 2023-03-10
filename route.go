package calculate

import (
	controller "calculate/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	main := r.Group("/calculate")
	{
		main.POST("/area", controller.CalArea())
		main.POST("/convert", controller.ConvertValue())
		main.POST("/convert/string", controller.ConvertValueString())
	}
}
