package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type EventName string

const (
	EventNameProductPassed   EventName = "product_passed"
	EventNameProductRejected EventName = "product_rejected"
	EventNameDeliveryCreated EventName = "delivery_created"
)

type Event struct {
	EventName string
	Time      time.Time
	Payload   interface{}
}

func (e *Event) Value() (driver.Value, error) {
	return json.Marshal(e)
}

func (e *Event) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		err := json.Unmarshal(data, &e)
		return err
	}
	return fmt.Errorf("failed to unmarshal subplier data")
}
