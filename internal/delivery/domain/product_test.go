package domain_test

import (
	"testing"

	"github.com/rattapon001/inventory-ddd/internal/delivery/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	assert := assert.New(t)

	product := domain.NewProduct("product-1", "sku-1", 100)

	assert.NotNil(product)
	assert.Equal("product-1", product.Name)
	assert.Equal("sku-1", product.Sku)
	assert.Equal(int16(100), product.Qty)
	assert.Equal(domain.ProductStatusPending, product.Status)
}

func TestPassProduct(t *testing.T) {
	assert := assert.New(t)

	product := domain.NewProduct("product-1", "sku-1", 100)

	err := product.Pass()
	assert.Nil(err)
	assert.Equal(domain.ProductStatusPass, product.Status)
}

func TestPassProduct_AlreadyPass(t *testing.T) {
	assert := assert.New(t)

	product := domain.NewProduct("product-1", "sku-1", 100)
	product.Pass()

	err := product.Pass()
	assert.NotNil(err)
	assert.Equal(domain.ProductStatusPass, product.Status)
}

func TestRejectProduct(t *testing.T) {
	assert := assert.New(t)

	product := domain.NewProduct("product-1", "sku-1", 100)

	err := product.Reject()
	assert.Nil(err)
	assert.Equal(domain.ProductStatusReject, product.Status)
}

func TestRejectProduct_AlreadyReject(t *testing.T) {
	assert := assert.New(t)

	product := domain.NewProduct("product-1", "sku-1", 100)
	product.Reject()

	err := product.Reject()
	assert.NotNil(err)
	assert.Equal(domain.ProductStatusReject, product.Status)
}
