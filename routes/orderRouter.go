package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/lolzni1/restaurant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("order/:order_id", controller.UpdateOrder())
}
