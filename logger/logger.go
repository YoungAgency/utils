package logger

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type CtxKey int

const (
	CtxRequestIdKey CtxKey = iota
	CtxErrKey       CtxKey = iota
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

func GinZerologMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a request id and add it to context and response headers
		requestID := uuid.New().String()
		ctx := c.Request.Context()
		valueCtx := context.WithValue(ctx, CtxRequestIdKey, requestID)
		c.Request = c.Request.WithContext(valueCtx)
		c.Writer.Header().Add("Request-Id", requestID)

		start := time.Now()
		c.Next()
		end := time.Since(start)

		event := logger.Log()
		event.
			Str("protocol", "http").
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("route", c.FullPath()).
			Str("ip", c.ClientIP()).
			Int("status", c.Writer.Status()).
			Dur("latency", end)

		err := c.Request.Context().Value(CtxErrKey)
		if err != nil {
			err, ok := err.(error)
			if ok {
				event.Err(err)
			}
		}

		// Emit log
		event.Msg("")
	}
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
