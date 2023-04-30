package service

import (
	"context"
	"fmt"
	"github.com/igorok-follow/analytics-server/app/models"
	context_tools "github.com/igorok-follow/analytics-server/tools/context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"strconv"
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
	span := trace.SpanFromContext(ctx)
	// разобраться с object атрибутами
	span.AddEvent("Service: "+context_tools.ExtractGRPCFullMethod(ctx), trace.WithAttributes(
		attribute.String("type", in.EventType),
		attribute.String("time", strconv.FormatInt(in.Unix, 10)),
	))
	defer span.End()

	log.Println(time.Now().UTC().Format("2006-01-02 03:04:05"))
	fmt.Println("grpc method:", context_tools.ExtractGRPCFullMethod(ctx), "request:", in)
	
	e.deps.EventHandler.Add(in)
	//if err != nil {
	//	err := errors.New("test internal error")
	//	span.RecordError(err)
	//	span.SetStatus(codes.Error, err.Error())
	//}

	return nil
}
