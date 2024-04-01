package delivery_router

import (
	"github.com/gin-gonic/gin"
	delivery_handler "github.com/rattapon001/inventory-ddd/api/handlers/delivery"
	"github.com/rattapon001/inventory-ddd/internal/delivery/app"
)

func InitRouter(router *gin.Engine, deliveryUseCase *app.DeliveryUseCase) {

	deliveryHandler := delivery_handler.NewDeliveryHandler(deliveryUseCase)

	delivery := router.Group("/delivery")
	{
		delivery.POST("/create", deliveryHandler.CreateDeliveryInfo)
	}
}
