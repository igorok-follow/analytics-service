package metadata

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const (
	RemoteAddress = "X-Remote-Address"
)

func ExtractRemoteAddress(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	remoteAddress := md.Get(RemoteAddress)
	if len(remoteAddress) == 0 {
		return ""
	}

	return remoteAddress[0]
}
