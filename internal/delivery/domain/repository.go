package domain

type DeliveryRepository interface {
	Save(deliveryInfo *DeliveryInformation) error
	GetByID(id DeliveryId) (*DeliveryInformation, error)
	GetAll() ([]*DeliveryInformation, error)
}
