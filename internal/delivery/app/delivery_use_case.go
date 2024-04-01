package app

import "github.com/rattapon001/inventory-ddd/internal/delivery/domain"

type DeliveryUseCase interface {
	Create(supplier domain.Supplier, products []domain.Product) (*domain.DeliveryInformation, error)
	Pass(deliveryId domain.DeliveryId, sku string) error
	Reject(deliveryId domain.DeliveryId, sku string) error
}

type deliveryUseCase struct {
	repo domain.DeliveryRepository
}

func NewDeliveryUseCase(repo domain.DeliveryRepository) DeliveryUseCase {
	return &deliveryUseCase{
		repo: repo,
	}
}
