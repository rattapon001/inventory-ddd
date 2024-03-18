package domain_test

import (
	"testing"

	"github.com/rattapon001/inventory-ddd/internal/inventory/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewReserve(t *testing.T) {
	assert := assert.New(t)

	reserve, err := domain.NewReserve("inventory-1", 100, "ref-1")

	assert.Nil(err)
	assert.NotNil(reserve)
	assert.NotEmpty(reserve.Id)
	assert.Equal(domain.InventoryId("inventory-1"), reserve.InventoryId)
	assert.Equal(100, reserve.Qty)
	assert.NotEmpty(reserve.ReservedAt)
	assert.Equal(domain.ReserveStatusPending, reserve.Status)
}

func TestReserveReserve(t *testing.T) {
	assert := assert.New(t)

	reserve, _ := domain.NewReserve("inventory-1", 100, "ref-1")

	err := reserve.Reserve()
	assert.Nil(err)
	assert.Equal(domain.ReserveStatusReserved, reserve.Status)
}

func TestReserveComplete(t *testing.T) {
	assert := assert.New(t)

	reserve, _ := domain.NewReserve("inventory-1", 100, "ref-1")

	err := reserve.Complete()
	assert.Nil(err)
	assert.Equal(domain.ReserveStatusComplete, reserve.Status)
}

func TestReserveCancel(t *testing.T) {
	assert := assert.New(t)

	reserve, _ := domain.NewReserve("inventory-1", 100, "ref-1")

	err := reserve.Cancel()
	assert.Nil(err)
	assert.Equal(domain.ReserveStatusCancel, reserve.Status)
}
