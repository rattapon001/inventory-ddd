package app

import "github.com/rattapon001/inventory-ddd/internal/delivery/domain"

func (uc *deliveryUseCase) Pass(deliveryId domain.DeliveryId, sku string) error {
	deliveryInfo, err := uc.repo.GetByID(deliveryId)
	if err != nil {
		return err
	}

	err = deliveryInfo.Pass(sku)
	if err != nil {
		return err
	}

	return uc.repo.Save(deliveryInfo)
}
