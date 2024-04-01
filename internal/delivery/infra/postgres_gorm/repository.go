package delivery_gorm

import (
	"github.com/rattapon001/inventory-ddd/internal/delivery/domain"
	"gorm.io/gorm"
)

type DeliveryGormRepository struct {
	DB *gorm.DB
}

func NewDeliveryGormRepository(db *gorm.DB) domain.DeliveryRepository {
	return &DeliveryGormRepository{
		DB: db,
	}
}

func (r *DeliveryGormRepository) Save(deliveryInfo *domain.DeliveryInformation) error {
	var existingDelivery domain.DeliveryInformation
	if err := r.DB.Where("id = ?", deliveryInfo.ID).First(&existingDelivery).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.DB.Create(deliveryInfo).Error
		}
		return err
	}
	return r.DB.Model(&existingDelivery).Updates(deliveryInfo).Error
}

func (r *DeliveryGormRepository) GetByID(id domain.DeliveryId) (*domain.DeliveryInformation, error) {
	var delivery domain.DeliveryInformation
	if err := r.DB.Where("id = ?", id).First(&delivery).Error; err != nil {
		return nil, err
	}
	return &delivery, nil
}

func (r *DeliveryGormRepository) GetAll() ([]*domain.DeliveryInformation, error) {
	var deliveries []*domain.DeliveryInformation
	if err := r.DB.Find(&deliveries).Error; err != nil {
		return nil, err
	}
	return deliveries, nil
}
