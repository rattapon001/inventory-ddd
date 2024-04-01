package app

import "github.com/rattapon001/inventory-ddd/internal/delivery/domain"

func (uc *deliveryUseCase) Reject(deliveryId domain.DeliveryId, sku string) error {
	deliveryInfo, err := uc.repo.GetByID(deliveryId)
	if err != nil {
		return err
	}

	err = deliveryInfo.Reject(sku)
	if err != nil {
		return err
	}

	return uc.repo.Save(deliveryInfo)
}
