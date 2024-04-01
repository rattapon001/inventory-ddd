package delivery_memory

import "github.com/rattapon001/inventory-ddd/internal/delivery/domain"

type DeliveryMemoryRepository struct {
	Deliveries []*domain.DeliveryInformation
}

func NewDeliveryMemoryRepository() *DeliveryMemoryRepository {
	return &DeliveryMemoryRepository{}
}

func (r *DeliveryMemoryRepository) Save(deliveryInfo *domain.DeliveryInformation) error {
	for i, delivery := range r.Deliveries {
		if delivery.ID == deliveryInfo.ID {
			r.Deliveries[i] = deliveryInfo
			return nil
		}
	}
	r.Deliveries = append(r.Deliveries, deliveryInfo)
	return nil
}

func (r *DeliveryMemoryRepository) GetByID(id domain.DeliveryId) (*domain.DeliveryInformation, error) {
	for _, delivery := range r.Deliveries {
		if delivery.ID == id {
			return delivery, nil
		}
	}
	return nil, nil
}

func (r *DeliveryMemoryRepository) GetAll() ([]*domain.DeliveryInformation, error) {
	return r.Deliveries, nil
}
