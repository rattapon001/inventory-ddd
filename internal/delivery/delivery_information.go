package delivery

import "time"

type DeliveryId string
type ProductStatus string

const (
	ProductStatusPending = "pending"
	ProductStatusPass    = "pass"
	ProductStatusReject  = "reject"
)

type Product struct {
	Name   string
	Sku    string
	Qty    int16
	Status ProductStatus
}

type Subplier struct {
	Name string
	Code string
}

type DeliveryInformation struct {
	ID             DeliveryId
	DocumentNumber string
	Date           time.Time
	Subplier       Subplier
	Products       []Product
}

func New(product []Product, subplier Subplier) (DeliveryInformation, error) {
	return DeliveryInformation{
		ID:             DeliveryId(time.Now().Format("mmddyyhhMMss")),
		DocumentNumber: time.Now().Format("mmddyyhhMMss"),
		Subplier:       subplier,
		Products:       product,
		Date:           time.Now(),
	}, nil
}

func (d *DeliveryInformation) Pass(sku string) error {
	for i, product := range d.Products {
		if product.Sku == sku {
			d.Products[i].Status = ProductStatusPass
		}
	}
	return nil
}

func (d *DeliveryInformation) Reject(sku string) error {
	for i, product := range d.Products {
		if product.Sku == sku {
			d.Products[i].Status = ProductStatusReject
		}
	}
	return nil
}
