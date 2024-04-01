package app

import "github.com/rattapon001/inventory-ddd/internal/delivery/domain"

func (uc *deliveryUseCase) Create(supplier domain.Supplier, products []domain.Product) (*domain.DeliveryInformation, error) {
	deliveryInfo, err := domain.New(products, supplier)

	if err != nil {
		return nil, err
	}

	err = uc.repo.Save(deliveryInfo)
	if err != nil {
		return nil, err
	}

	return deliveryInfo, nil
}
