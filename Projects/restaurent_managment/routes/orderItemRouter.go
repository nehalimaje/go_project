package routes

import (
	controllers "golang-restaurent-managment/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItemid", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:orderid", controllers.GetOrderItemByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItemid", controllers.UpdateOrderItem())
}
