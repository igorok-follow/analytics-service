package service

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/igorok-follow/analytics-service/app/models"
	"github.com/igorok-follow/analytics-service/app/repository"
	"github.com/igorok-follow/analytics-service/tools/event_handler"
	"go.opentelemetry.io/otel/trace"
)

type Container struct {
	EventService EventService
}

func NewServices(deps *Dependencies) *Container {
	eventService := NewEventService(deps)

	return &Container{
		EventService: eventService,
	}
}

type EventService interface {
	RegisterEvent(ctx context.Context, in *models.Event) error
}

type Dependencies struct {
	Redis        *redis.Client
	EventHandler *event_handler.EventHandler
	Tracer       trace.Tracer
	Repositories *repository.Container
}
