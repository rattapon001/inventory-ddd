package domain_test

import (
	"testing"
	"time"

	"github.com/rattapon001/inventory-ddd/internal/inventory/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewInventory(t *testing.T) {
	assert := assert.New(t)

	inventory, err := domain.NewInventory("lot-1", 100, time.Now())

	assert.Nil(err)
	assert.NotNil(inventory)
	assert.NotEmpty(inventory.Id)
	assert.Equal("lot-1", inventory.LotNumber)
	assert.Equal(100, inventory.Qty)
	assert.NotEmpty(inventory.Eta)

}

func TestInventoryDeductQty(t *testing.T) {
	assert := assert.New(t)

	inventory, _ := domain.NewInventory("lot-1", 100, time.Now())

	err := inventory.DeductQty(50)
	assert.Nil(err)
	assert.Equal(50, inventory.Qty)
}

func TestInventoryDeductQtyNotEnough(t *testing.T) {
	assert := assert.New(t)

	inventory, _ := domain.NewInventory("lot-1", 100, time.Now())

	err := inventory.DeductQty(150)
	assert.NotNil(err)
	assert.Equal("inventory qty not enough", err.Error())
}
