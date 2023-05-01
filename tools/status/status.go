package status

import (
	"context"
	"github.com/igorok-follow/analytics-server/tools/tracing"
	otel_codes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpc_codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewError(ctx context.Context, tr trace.Tracer, err error, code grpc_codes.Code) error {
	ctx, span := tracing.SpanFromContext(
		ctx,
		tr,
		"status",
		"status.Error",
	)
	defer span.End()

	span.RecordError(err)
	span.SetStatus(otel_codes.Error, "")

	return status.Error(code, err.Error())
}

//docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 -e COLLECTOR_OTLP_ENABLED=true -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 4317:4317 -p 4318:4318 -p 14250:14250 -p 14268:14268 -p 14269:14269 -p 9411:9411 --restart unless-stopped jaegertracing/all-in-one:1.44
