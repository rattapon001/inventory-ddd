package domain_test

import (
	"testing"
	"time"

	inventory "github.com/rattapon001/inventory-ddd/internal/inventory/domain"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	product, err := inventory.New("sku", "name")

	assert.Nil(err)
	assert.NotNil(product)
	assert.NotEmpty(product.Id)
	assert.Equal("sku", product.Sku)
	assert.Equal("name", product.Name)
}

func TestDeposit(t *testing.T) {
	assert := assert.New(t)

	product, err := inventory.New("sku", "name")
	assert.Nil(err)

	eta := time.Now()

	err = product.Deposit("lot-1", 100, eta)
	assert.Nil(err)
	assert.Len(product.Inventory, 1)
	assert.Equal("lot-1", product.Inventory[0].LotNumber)
	assert.Equal(100, product.Inventory[0].Qty)
	assert.Equal(eta, product.Inventory[0].Eta)
}

func TestSortInventoryByETA(t *testing.T) {
	assert := assert.New(t)

	product, err := inventory.New("sku", "name")
	assert.Nil(err)

	eta1 := time.Now()
	eta2 := time.Now().AddDate(0, 0, 1)

	err = product.Deposit("lot-2", 100, eta2)
	assert.Nil(err)

	err = product.Deposit("lot-1", 100, eta1)
	assert.Nil(err)

	assert.Len(product.Inventory, 2)
	assert.Equal(eta2, product.Inventory[0].Eta)
	assert.Equal(eta1, product.Inventory[1].Eta)

	product.SortInventoryByETA()
	assert.Equal(eta1, product.Inventory[0].Eta)
	assert.Equal(eta2, product.Inventory[1].Eta)
}

func TestReserveInventory(t *testing.T) {
	assert := assert.New(t)

	product, err := inventory.New("sku", "name")
	assert.Nil(err)

	eta := time.Now()

	err = product.Deposit("lot-1", 100, eta)
	assert.Nil(err)

	err = product.ReserveInventory("lot-1", 50, "ref-1")
	assert.Nil(err)
	assert.Len(product.Inventory, 1)
	assert.Len(product.Reserve, 1)
	assert.Equal("lot-1", product.Reserve[0].LotNumber)
	assert.Equal(50, product.Reserve[0].Qty)
	assert.Equal("ref-1", product.Reserve[0].RefNo)
}
