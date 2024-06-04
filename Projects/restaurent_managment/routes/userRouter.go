package routes

import (
	"golang-restaurent-managment/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:userid", controllers.GetUser())
	incomingRoutes.POST("/signup", controllers.SignUp())
	incomingRoutes.POST("/login", controllers.Login())
}
