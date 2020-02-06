package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// logDelta writes the provided message to the log along with how long the RPC took too complete.
func logDelta(start time.Time, message string) {
	log.Printf("%s (took %s)", message, time.Since(start))
}

// gRPC provides middleware hooks for "Unary" RPCs (i.e. those that follow the simpler request/response)
// model and for "Streaming" RPCs. The signature differs slightly due to the way that gRPC works internally
// but the grpc.ServerStream parameter provides access to the request Context and thus the middleware
// that stores the RPC metadata (in this case the authorization headers).
func UnaryLoggingMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("couldn't parse metadata")
	}
	defer logDelta(time.Now(), fmt.Sprintf("user request: %s called %s", md.Get("authorization"), info.FullMethod))
	return handler(ctx, req)
}

func StreamingLoggingMiddleware(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errors.New("couldn't parse metadata")
	}
	defer logDelta(time.Now(), fmt.Sprintf("user request: %s called %s", md.Get("authorization"), info.FullMethod))
	return handler(srv, ss)
}
