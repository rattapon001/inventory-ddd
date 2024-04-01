package domain

import (
	"time"

	"github.com/google/uuid"
	errs "github.com/rattapon001/inventory-ddd/internal/delivery/errors"
)

type DeliveryId string

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
	Events         []Event
}

func New(product []Product, subplier Subplier) (*DeliveryInformation, error) {

	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	deliveryInformation := &DeliveryInformation{
		ID:             DeliveryId(ID.String()),
		DocumentNumber: time.Now().Format("mmddyyhhMMss"),
		Subplier:       subplier,
		Products:       product,
		Date:           time.Now(),
	}

	deliveryInformation.DeliveryInformationCreatedEvent()

	return deliveryInformation, nil
}

func (d *DeliveryInformation) Pass(sku string) error {
	for i, product := range d.Products {
		if product.Sku == sku {
			err := d.Products[i].Pass()
			if err != nil {
				return err
			}
			d.ProductPassedEvent(d.Products[i])
			return nil
		}
	}
	return errs.ErrProductNotFound
}

func (d *DeliveryInformation) Reject(sku string) error {
	for i, product := range d.Products {
		if product.Sku == sku {
			err := d.Products[i].Reject()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errs.ErrProductNotFound
}

func (d *DeliveryInformation) ProductPassedEvent(product Product) {

	event := Event{
		EventName: string(EventNameProductPassed),
		Time:      time.Now(),
		Payload: map[string]interface{}{
			"Subplier": d.Subplier,
			"Product":  product,
			"Eta":      d.Date,
		},
	}

	d.Events = append(d.Events, event)
}

func (d *DeliveryInformation) ProductRejectedEvent(product Product) {

	event := Event{
		EventName: string(EventNameProductRejected),
		Time:      time.Now(),
		Payload: map[string]interface{}{
			"Subplier": d.Subplier,
			"Product":  product,
			"Eta":      d.Date,
		},
	}

	d.Events = append(d.Events, event)
}

func (d *DeliveryInformation) DeliveryInformationCreatedEvent() {

	event := Event{
		EventName: string(EventNameDeliveryCreated),
		Time:      time.Now(),
		Payload: map[string]interface{}{
			"Id":             d.ID,
			"DocumentNumber": d.DocumentNumber,
			"Date":           d.Date,
		},
	}

	d.Events = append(d.Events, event)
}
