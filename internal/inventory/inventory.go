package inventory

import (
	"time"

	"github.com/google/uuid"
)

type InventoryId string
type ProductId string
type ReserveId string
type ReserveStatus string

const (
	ReserveStatusPending  ReserveStatus = "pending"
	ReserveStatusComplete ReserveStatus = "completed"
	ReserveStatusCancel   ReserveStatus = "cancel"
)

type Product struct {
	Id        ProductId
	Sku       string
	Name      string
	Inventory []Inventory
	Reserve   Reserve
}

type Inventory struct {
	Id        InventoryId
	LotNumber string
	Qty       int
	Eta       time.Time
}

type Reserve struct {
	Id          ReserveId
	InventoryId InventoryId
	Qty         int
	RefNo       string
	Status      ReserveStatus
}

func New(sku string, name string) (*Product, error) {
	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Product{
		Id:   ProductId(ID.String()),
		Sku:  sku,
		Name: name,
	}, nil

}

func (p *Product) Deposit(lotNumber string, qty int, eta time.Time) error {
	ID, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	p.Inventory = append(p.Inventory, Inventory{
		Id:        InventoryId(ID.String()),
		LotNumber: lotNumber,
		Qty:       qty,
		Eta:       eta,
	})

	return nil
}
