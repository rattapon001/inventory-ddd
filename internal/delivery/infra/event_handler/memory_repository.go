package eventhandler

type MemoryOutboxRepository struct {
	outboxEvents []*OutboxEvent
}

func NewMemoryOutboxRepository() *MemoryOutboxRepository {
	return &MemoryOutboxRepository{
		outboxEvents: []*OutboxEvent{},
	}
}

func (r *MemoryOutboxRepository) Save(outboxEvent *OutboxEvent) error {
	r.outboxEvents = append(r.outboxEvents, outboxEvent)
	return nil
}
