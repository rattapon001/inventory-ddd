package domain

import (
	"time"

	"github.com/google/uuid"
)

type ReserveId string
type ReserveStatus string

const (
	ReserveStatusPending  ReserveStatus = "pending"
	ReserveStatusReserved ReserveStatus = "reserved"
	ReserveStatusComplete ReserveStatus = "completed"
	ReserveStatusCancel   ReserveStatus = "cancel"
)

type Reserve struct {
	Id          ReserveId
	ReservedAt  time.Time
	InventoryId InventoryId
	Qty         int
	RefNo       string
	Status      ReserveStatus
}

func NewReserve(inventoryId InventoryId, qty int, refNo string) (*Reserve, error) {
	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Reserve{
		Id:          ReserveId(ID.String()),
		ReservedAt:  time.Now(),
		InventoryId: inventoryId,
		Qty:         qty,
		RefNo:       refNo,
		Status:      ReserveStatusPending,
	}, nil
}

func (r *Reserve) Reserve() error {
	r.Status = ReserveStatusReserved
	return nil
}

func (r *Reserve) Complete() error {
	r.Status = ReserveStatusComplete
	return nil
}

func (r *Reserve) Cancel() error {
	r.Status = ReserveStatusCancel
	return nil
}
