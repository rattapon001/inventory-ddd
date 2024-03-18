package domain

import "time"

type ReserveId string
type ReserveStatus string

const (
	ReserveStatusPending  ReserveStatus = "pending"
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
