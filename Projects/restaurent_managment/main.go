package main

import (
	"golang-restaurent-managment/database"
	"golang-restaurent-managment/middleware"
	"golang-restaurent-managment/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection()

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
