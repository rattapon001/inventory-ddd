package eventhandler

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Publisher struct {
	repo OutboxRepository
}

func NewPublisher(db *gorm.DB) *Publisher {
	return &Publisher{
		repo: NewPostgresOutboxRepository(db),
	}
}

func (p *Publisher) Publish(evenName string, payload interface{}) error {

	aggregateID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	outboxEvent := NewOutboxEvent(aggregateID.String(), evenName, evenName, payload)
	return p.repo.Save(outboxEvent)
}
