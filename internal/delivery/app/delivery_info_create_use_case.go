package app

import "github.com/rattapon001/inventory-ddd/internal/delivery/domain"

func (uc *deliveryUseCase) Create(subplier domain.Subplier, products []domain.Product) (*domain.DeliveryInformation, error) {
	deliveryInfo, err := domain.New(products, subplier)

	if err != nil {
		return nil, err
	}

	err = uc.repo.Save(deliveryInfo)
	if err != nil {
		return nil, err
	}

	return deliveryInfo, nil
}
