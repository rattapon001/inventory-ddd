package domain_test

import (
	"testing"

	"github.com/rattapon001/inventory-ddd/internal/delivery/domain"
	errs "github.com/rattapon001/inventory-ddd/internal/delivery/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	delivery, err := domain.New([]domain.Product{
		{
			Name:   "product-1",
			Sku:    "sku-1",
			Qty:    100,
			Status: domain.ProductStatusPending,
		},
	}, domain.Supplier{
		Name: "supplier-1",
		Code: "code-1",
	})

	assert.Nil(err)
	assert.NotNil(delivery)
	assert.Equal("supplier-1", delivery.Supplier.Name)
	assert.Equal("code-1", delivery.Supplier.Code)
	assert.Equal(1, len(delivery.Products))
	assert.Equal("product-1", delivery.Products[0].Name)
	assert.Equal("sku-1", delivery.Products[0].Sku)
	assert.Equal(1, len(delivery.Events))
}

func TestPass(t *testing.T) {
	assert := assert.New(t)

	delivery, _ := domain.New([]domain.Product{
		{
			Name:   "product-1",
			Sku:    "sku-1",
			Qty:    100,
			Status: domain.ProductStatusPending,
		},
	}, domain.Supplier{
		Name: "supplier-1",
		Code: "code-1",
	})

	err := delivery.Pass("sku-1")
	assert.Nil(err)
	assert.Equal(domain.ProductStatusPass, delivery.Products[0].Status)
}

func TestPass_NotFound(t *testing.T) {
	assert := assert.New(t)

	delivery, _ := domain.New([]domain.Product{
		{
			Name:   "product-1",
			Sku:    "sku-1",
			Qty:    100,
			Status: domain.ProductStatusPending,
		},
	}, domain.Supplier{
		Name: "supplier-1",
		Code: "code-1",
	})

	err := delivery.Pass("sku-2")
	assert.NotNil(err)
	assert.Equal(errs.ErrProductNotFound, err)
}

func TestReject(t *testing.T) {
	assert := assert.New(t)

	delivery, _ := domain.New([]domain.Product{
		{
			Name:   "product-1",
			Sku:    "sku-1",
			Qty:    100,
			Status: domain.ProductStatusPending,
		},
	}, domain.Supplier{
		Name: "supplier-1",
		Code: "code-1",
	})

	err := delivery.Reject("sku-1")
	assert.Nil(err)
	assert.Equal(domain.ProductStatusReject, delivery.Products[0].Status)
}

func TestReject_NotFound(t *testing.T) {
	assert := assert.New(t)

	delivery, _ := domain.New([]domain.Product{
		{
			Name:   "product-1",
			Sku:    "sku-1",
			Qty:    100,
			Status: domain.ProductStatusPending,
		},
	}, domain.Supplier{
		Name: "supplier-1",
		Code: "code-1",
	})

	err := delivery.Reject("sku-2")
	assert.NotNil(err)
	assert.Equal(errs.ErrProductNotFound, err)
}

func TestProductPassedEvent(t *testing.T) {
	assert := assert.New(t)

	delivery, _ := domain.New([]domain.Product{
		{
			Name:   "product-1",
			Sku:    "sku-1",
			Qty:    100,
			Status: domain.ProductStatusPending,
		},
	}, domain.Supplier{
		Name: "supplier-1",
		Code: "code-1",
	})

	delivery.ProductPassedEvent(delivery.Products[0])
	assert.Equal(2, len(delivery.Events))
}

func TestProductRejectedEvent(t *testing.T) {
	assert := assert.New(t)

	delivery, _ := domain.New([]domain.Product{
		{
			Name:   "product-1",
			Sku:    "sku-1",
			Qty:    100,
			Status: domain.ProductStatusPending,
		},
	}, domain.Supplier{
		Name: "supplier-1",
		Code: "code-1",
	})

	delivery.ProductRejectedEvent(delivery.Products[0])
	assert.Equal(2, len(delivery.Events))
}
