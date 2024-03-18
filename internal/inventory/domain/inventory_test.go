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
