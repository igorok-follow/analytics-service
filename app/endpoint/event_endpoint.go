package endpoint

import (
	"context"
	"github.com/igorok-follow/analytics-service/app/models"
	"github.com/igorok-follow/analytics-service/app/service"
	"github.com/igorok-follow/analytics-service/extra/api"
	context_tools "github.com/igorok-follow/analytics-service/tools/context"
	"github.com/igorok-follow/analytics-service/tools/metadata"
	status_tools "github.com/igorok-follow/analytics-service/tools/status"
	"github.com/igorok-follow/analytics-service/tools/tracing"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
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
	var span trace.Span
	ctx, span = tracing.SpanFromContext(
		ctx,
		e.Tracer,
		"endpoint",
		context_tools.ExtractGRPCFullMethod(ctx),
	)
	defer span.End()

	err := e.Services.EventService.RegisterEvent(ctx, &models.Event{
		UserId:    metadata.ExtractRemoteAddress(ctx),
		EventType: in.EventType,
		Unix:      time.Now().Unix(),
	})
	if err != nil {
		return nil, status_tools.NewError(ctx, e.Tracer, err, codes.Internal)
	}

	return &api.Empty{}, nil
}
