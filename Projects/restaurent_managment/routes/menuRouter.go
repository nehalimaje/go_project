package routes

import (
	controllers "golang-restaurent-managment/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controllers.GetMenus())
	incomingRoutes.GET("/menu/:menuid", controllers.GetMenu())
	incomingRoutes.POST("/menu", controllers.CreateMenu())
	incomingRoutes.PATCH("/menu/:menuid", controllers.UpdateMenu())
}
