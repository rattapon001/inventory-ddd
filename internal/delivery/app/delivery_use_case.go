package app

import "github.com/rattapon001/inventory-ddd/internal/delivery/domain"

type DeliveryUseCase interface {
	Create(documentNumber string, subplier domain.Subplier, products []domain.Product) (*domain.DeliveryInformation, error)
	Pass(deliveryId domain.DeliveryId, sku string) error
	Reject(deliveryId domain.DeliveryId, sku string) error
}
