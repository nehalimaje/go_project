package routes

import (
	controllers "golang-restaurent-managment/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controllers.GetOrders())
	incomingRoutes.GET("/order/:orderid", controllers.GetOrder())
	incomingRoutes.POST("/order", controllers.CreateOrder())
	incomingRoutes.PATCH("/order/:orderid", controllers.UpdateOrder())
}
