package eventhandler

import "gorm.io/gorm"

type OutboxRepository interface {
	Save(outboxEvent *OutboxEvent) error
}

type PostgresOutboxRepository struct {
	db *gorm.DB
}

func NewPostgresOutboxRepository(db *gorm.DB) *PostgresOutboxRepository {
	return &PostgresOutboxRepository{
		db: db,
	}
}

func (r *PostgresOutboxRepository) Save(outboxEvent *OutboxEvent) error {
	return r.db.Create(outboxEvent).Error
}
