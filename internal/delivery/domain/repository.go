package domain

type Repository interface {
	save(deliveryInfo *DeliveryInformation) error
	getByID(id DeliveryId) (*DeliveryInformation, error)
	getAll() ([]*DeliveryInformation, error)
}
