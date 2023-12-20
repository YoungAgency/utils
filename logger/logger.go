package logger

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

// NewZerologLogger returns an instance of a zerolog Logger.
// It can use two formats: json and text. The service parameter
// is logged in each log line.
func NewZerologLogger(service string, format string, out io.Writer) *zerolog.Logger {
	zerolog.TimestampFieldName = "timestamp"
	var output io.Writer
	switch format {
	case "test":
		output = zerolog.ConsoleWriter{Out: out, TimeFormat: time.RFC3339}
	case "json":
		output = out
	default:
		panic(errors.New("unsupported format"))
	}
	l := zerolog.New(output).With().Timestamp().Str("service", service).Logger()
	return &l
}

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
