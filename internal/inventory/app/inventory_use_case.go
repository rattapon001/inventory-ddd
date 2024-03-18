package app

import (
	"time"

	"github.com/rattapon001/inventory-ddd/internal/delivery/domain"
)

type InventoryUseCase interface {
	// Deposit product to inventory
	Deposit(sku string, name string, qty int16, eta time.Time) (*domain.Product, error)
}
