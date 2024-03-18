package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProductId string

type Product struct {
	Id        ProductId
	Sku       string
	Name      string
	Inventory []Inventory
	Reserve   []Reserve
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

	Inventory, err := NewInventory(lotNumber, qty, eta)

	if err != nil {
		return err
	}

	p.Inventory = append(p.Inventory, *Inventory)

	return nil
}

// Sort By ETA Ascending
func (p *Product) SortInventoryByETA() {

	for i := 0; i < len(p.Inventory); i++ {
		for j := i + 1; j < len(p.Inventory); j++ {
			if p.Inventory[i].Eta.After(p.Inventory[j].Eta) {
				p.Inventory[i], p.Inventory[j] = p.Inventory[j], p.Inventory[i]
			}
		}
	}

}
