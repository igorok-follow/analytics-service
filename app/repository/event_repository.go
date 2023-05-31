package repository

import (
	"context"
	"github.com/igorok-follow/analytics-service/app/models"
	"github.com/jmoiron/sqlx"
)

type Event struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *Event {
	return &Event{
		db: db,
	}
}

func (r *Event) InsertEvent(ctx context.Context, e *models.Event) error {
	_, err := r.db.ExecContext(ctx, "insert into events(user_id, event_type, created) values ($1, $2, $3)", e.UserId, e.EventType, e.Unix)
	if err != nil {
		return err
	}

	return nil
}
