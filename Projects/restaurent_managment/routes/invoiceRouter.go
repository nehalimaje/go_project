package routes

import (
	controllers "golang-restaurent-managment/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoice", controllers.GetInvoices())
	incomingRoutes.GET("/invoice/:invoiceid", controllers.GetInvoice())
	incomingRoutes.POST("/invoice", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoice/:invoiceid", controllers.UpdateInvoice())
}
