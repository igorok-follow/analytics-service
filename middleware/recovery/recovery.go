package recovery

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				err = errors.New("internal server error")
			}
		}()

		resp, err = handler(ctx, req)

		return resp, err
	}
}
