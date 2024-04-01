package delivery_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/inventory-ddd/internal/delivery/app"
	"github.com/rattapon001/inventory-ddd/internal/delivery/domain"
)

type DeliveryHandler struct {
	deliveryUseCase app.DeliveryUseCase
}

func NewDeliveryHandler(deliveryUseCase *app.DeliveryUseCase) DeliveryHandler {
	return DeliveryHandler{
		deliveryUseCase: *deliveryUseCase,
	}
}

func (h *DeliveryHandler) CreateDeliveryInfo(c *gin.Context) {
	var deliveryInfo domain.DeliveryInformation
	if err := c.BindJSON(&deliveryInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delivery, err := h.deliveryUseCase.Create(deliveryInfo.Supplier, deliveryInfo.Products)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, delivery)
}
