package service

import (
	"context"
	"fmt"
	"github.com/igorok-follow/analytics-service/app/models"
	context_tools "github.com/igorok-follow/analytics-service/tools/context"
	"github.com/igorok-follow/analytics-service/tools/metadata"
	"github.com/igorok-follow/analytics-service/tools/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"time"
)

type Event struct {
	deps *Dependencies
}

func NewEventService(
	deps *Dependencies,
) *Event {
	return &Event{
		deps: deps,
	}
}

func (e *Event) RegisterEvent(ctx context.Context, in *models.Event) error {
	var span trace.Span
	ctx, span = tracing.SpanFromContext(
		ctx,
		e.deps.Tracer,
		"service",
		context_tools.ExtractGRPCFullMethod(ctx),
		attribute.KeyValue{
			Key:   "event_type",
			Value: attribute.StringValue(in.EventType),
		},
		attribute.KeyValue{
			Key:   "ip",
			Value: attribute.StringValue(metadata.ExtractRemoteAddress(ctx)),
		},
	)
	defer span.End()

	log.Println(time.Now().UTC().Format("2006-01-02 03:04:05"))
	fmt.Println("grpc method:", context_tools.ExtractGRPCFullMethod(ctx), "request:", in)

	err := e.deps.EventHandler.Add(in)
	if err != nil {
		return err
	}

	return nil
}
