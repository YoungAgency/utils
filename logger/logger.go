package logger

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func GrpcZerologMiddleware(logger *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// status
		route := info.FullMethod
		event := logger.Log()

		start := time.Now()
		resp, err := handler(ctx, req)
		end := time.Since(start)

		event.
			Str("protocol", "grpc").
			Str("route", route).
			Dur("latency", end).
			Err(err).
			Msg("")

		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
