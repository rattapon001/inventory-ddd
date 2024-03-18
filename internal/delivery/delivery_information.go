package delivery

import (
	"time"

	"github.com/google/uuid"
)

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

func New(product []Product, subplier Subplier) (*DeliveryInformation, error) {

	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &DeliveryInformation{
		ID:             DeliveryId(ID.String()),
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
