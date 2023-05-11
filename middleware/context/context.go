package context

import (
	"context"
	context_tools "github.com/igorok-follow/analytics-service/tools/context"
	"github.com/igorok-follow/analytics-service/tools/metadata"
	"google.golang.org/grpc"
	"path"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		interface{},
		error,
	) {
		ctx = ToContext(ctx, info.FullMethod)

		resp, err := handler(ctx, req)

		return resp, err
	}
}

func ToContext(
	ctx context.Context,
	fullMethod string,
) context.Context {
	ctx = FromMetadataToContext(ctx)
	ctx = GRPCIntoToContext(ctx, fullMethod)

	return ctx
}

func FromMetadataToContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, metadata.RemoteAddress, metadata.ExtractRemoteAddress(ctx))

	return ctx
}

func GRPCIntoToContext(ctx context.Context, fullMethod string) context.Context {
	fullMethod = path.Base(fullMethod)

	ctx = context.WithValue(ctx, context_tools.GRPCFullMethod, fullMethod)

	return ctx
}
