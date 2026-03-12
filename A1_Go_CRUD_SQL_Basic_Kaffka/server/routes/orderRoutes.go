package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {

	r.POST("/order", controllers.CreateOrders)
	r.GET("/order", controllers.GetOrders)
	r.GET("/order/:order_id", controllers.GetOrdersById)

}
