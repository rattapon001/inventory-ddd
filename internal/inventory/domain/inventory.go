package domain

import (
	"time"

	"github.com/google/uuid"
	errs "github.com/rattapon001/inventory-ddd/internal/inventory/errors"
)

type InventoryId string

type Inventory struct {
	Id        InventoryId
	LotNumber string
	Qty       int
	Eta       time.Time
}

func NewInventory(lotNumber string, qty int, eta time.Time) (*Inventory, error) {
	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err

	}
	return &Inventory{
		Id:        InventoryId(ID.String()),
		LotNumber: lotNumber,
		Qty:       qty,
		Eta:       eta,
	}, nil
}

func (i *Inventory) DeductQty(qty int) error {

	if i.Qty < qty {
		return errs.ErrInventoryQtyNotEnough
	}

	i.Qty = i.Qty - qty

	return nil
}
