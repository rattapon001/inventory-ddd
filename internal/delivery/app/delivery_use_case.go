package app

import (
	eventhandler "github.com/rattapon001/inventory-ddd/internal/delivery/app/event_handler"
	"github.com/rattapon001/inventory-ddd/internal/delivery/domain"
)

type DeliveryUseCase interface {
	Create(supplier domain.Supplier, products []domain.Product) (*domain.DeliveryInformation, error)
	Pass(deliveryId domain.DeliveryId, sku string) error
	Reject(deliveryId domain.DeliveryId, sku string) error
}

type deliveryUseCase struct {
	repo      domain.DeliveryRepository
	publisher eventhandler.EventHandler
}

func NewDeliveryUseCase(repo domain.DeliveryRepository, publisher eventhandler.EventHandler) DeliveryUseCase {
	return &deliveryUseCase{
		repo:      repo,
		publisher: publisher,
	}
}
