package routes

import (
	controllers "golang-restaurent-managment/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:foodid", controllers.GetFood())
	incomingRoutes.POST("/food", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:foodid", controllers.UpdateFood())
}
