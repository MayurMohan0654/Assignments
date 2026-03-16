package routes

import (
	"server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, cont controllers.OrderController) {

	r.POST("/order", cont.CreateOrders)
	r.GET("/order", cont.GetOrders)
	r.GET("/order/:order_id", cont.GetOrderById)

}
