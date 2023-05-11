package validate

import (
	"context"
	"github.com/igorok-follow/analytics-service/extra/api"
	"github.com/igorok-follow/analytics-service/helpers"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		interface{},
		error,
	) {
		err := Validate(req, info)
		if err != nil {
			return nil, err
		}

		resp, err := handler(ctx, req)
		return resp, err
	}
}

func Validate(req interface{}, info *grpc.UnaryServerInfo) error {
	switch info.FullMethod {
	case "store.Store/GetCatalog":
		err := helpers.ValidateRegisterEvent(req.(*api.RegisterEventReq))
		if err != nil {
			return err
		}
	}

	return nil
}
