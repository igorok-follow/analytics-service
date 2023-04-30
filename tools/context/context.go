package context

import (
	"context"
)

const (
	GRPCFullMethod = "GRPC_FULL_METHOD"
	SpanKey        = "SPAN_KEY"
)

func ExtractGRPCFullMethod(ctx context.Context) string {
	fullMethod, ok := ctx.Value(GRPCFullMethod).(string)
	if !ok {
		return ""
	}

	return fullMethod
}
