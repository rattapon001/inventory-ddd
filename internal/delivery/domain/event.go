package domain

import "time"

type Event struct {
	EventName string
	Time      time.Time
	Payload   interface{}
}

type EventName string

const (
	EventNameProductPassed   EventName = "product_passed"
	EventNameProductRejected EventName = "product_rejected"
	EventNameDeliveryCreated EventName = "delivery_created"
)
