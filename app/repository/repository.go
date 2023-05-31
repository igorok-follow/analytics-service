package repository

import (
	"context"
	"github.com/igorok-follow/analytics-service/app/models"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	EventRepository EventRepository
}

func NewRepositoryContainer(db *sqlx.DB) *Container {
	return &Container{
		EventRepository: NewEventRepository(db),
	}
}

type EventRepository interface {
	InsertEvent(ctx context.Context, e *models.Event) error
}
