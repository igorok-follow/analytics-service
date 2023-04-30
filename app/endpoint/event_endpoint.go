package endpoint

import (
	"context"
	"github.com/igorok-follow/analytics-server/app/models"
	"github.com/igorok-follow/analytics-server/app/service"
	"github.com/igorok-follow/analytics-server/extra/api"
	"github.com/igorok-follow/analytics-server/extra/error_codes"
	context_tools "github.com/igorok-follow/analytics-server/tools/context"
	"github.com/igorok-follow/analytics-server/tools/metadata"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Event struct {
	Services *service.Container
	Tracer   trace.Tracer
}

func NewEventEndpoint(services *service.Container, tracer trace.Tracer) *Event {
	return &Event{
		Services: services,
		Tracer:   tracer,
	}
}

func (e *Event) RegisterEvent(ctx context.Context, in *api.RegisterEventReq) (*api.Empty, error) {
	// validate...

	ctx, span := e.Tracer.Start(ctx, context_tools.ExtractGRPCFullMethod(ctx))
	span.AddEvent("Endpoint: "+context_tools.ExtractGRPCFullMethod(ctx), trace.WithAttributes(attribute.String("ip", metadata.ExtractRemoteAddress(ctx))))
	defer span.End()

	err := e.Services.EventService.RegisterEvent(ctx, &models.Event{
		UserId:    metadata.ExtractRemoteAddress(ctx),
		EventType: in.EventType,
		Unix:      time.Now().Unix(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, error_codes.INTERNAL)
	}

	return &api.Empty{}, nil
}
