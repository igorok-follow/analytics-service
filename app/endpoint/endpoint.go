package endpoint

import (
	"context"
	"github.com/igorok-follow/analytics-service/app/service"
	"github.com/igorok-follow/analytics-service/extra/api"
)

type Container struct {
	EventEndpoint EventEndpoint
}

func NewEndpointContainer(services *service.Container, deps *service.Dependencies) *Container {
	return &Container{
		EventEndpoint: NewEventEndpoint(services, deps.Tracer),
	}
}

type EventEndpoint interface {
	RegisterEvent(ctx context.Context, in *api.RegisterEventReq) (*api.Empty, error)
}
