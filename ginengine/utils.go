package ginengine

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type CtxKey int

const (
	CtxRequestIdKey CtxKey = iota
	CtxErrKey       CtxKey = iota
)

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

		var ip string

		if cloudflareIp := c.Request.Header.Get("CF-Connecting-IP"); cloudflareIp != "" {
			ip = cloudflareIp
		} else if forwardedIp := c.Request.Header.Get("X-Forwarded-For"); forwardedIp != "" {
			ip = forwardedIp
		} else {
			ip = c.ClientIP()
		}

		event := logger.Log()
		event.
			Str("protocol", "http").
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("route", c.FullPath()).
			Str("query", c.Request.URL.RawQuery).
			Str("ip", ip).
			Str("user_agent", c.Request.UserAgent()).
			Int("status", c.Writer.Status()).
			Dur("latency", end)

		// context key set by auth package
		if val, ok := c.Get("young-user"); val != nil && ok {
			event.Interface("user", val)
		}

		err := c.Request.Context().Value(CtxErrKey)
		if err != nil {
			err, ok := err.(error)
			if ok {
				event.Err(err)
			}
		}

		// Emit log
		event.Send()
	}
}
