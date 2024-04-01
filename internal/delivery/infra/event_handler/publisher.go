package eventhandler

import (
	"github.com/google/uuid"
)

type Publisher struct {
	repo OutboxRepository
}

func NewPublisher(repository OutboxRepository) *Publisher {
	return &Publisher{
		repo: repository,
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
