package routes

import (
	controllers "golang-restaurent-managment/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/table", controllers.GetTables())
	incomingRoutes.GET("/table/:tableid", controllers.GetTable())
	incomingRoutes.POST("/table", controllers.CreateTable())
	incomingRoutes.PATCH("/table/:tableid", controllers.UpdateTable())
}
