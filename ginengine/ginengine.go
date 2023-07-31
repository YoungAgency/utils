package ginengine

import (
	"time"

	"github.com/YoungAgency/rate"
	rateGin "github.com/YoungAgency/rate/gin"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type GinWrapper struct {
	engine *gin.Engine
}

func New() *GinWrapper {
	engine := gin.New()
	return &GinWrapper{
		engine: engine,
	}
}

// DefaultGinEngine returns a default gin engine with the following middlewares: recovery, sentry, logger, cors, rateLimiter and healthCheck
func DefaultGinEngine(allowedOrigins []string, logger *zerolog.Logger) *gin.Engine {
	engine := New().
		WithRecovery().
		WithSentry().
		WithZerolog(logger)

	if len(allowedOrigins) > 0 {
		engine = engine.WithCors(allowedOrigins)
	}

	return engine.WithRateLimiter().
		WithHealthCheck().
		Engine()
}

func (g *GinWrapper) WithRecovery() *GinWrapper {
	g.engine.Use(gin.Recovery())
	return g
}

func (g *GinWrapper) WithZerolog(logger *zerolog.Logger) *GinWrapper {
	g.engine.Use(GinZerologMiddleware(logger))
	return g
}

func (g *GinWrapper) WithSentry() *GinWrapper {
	g.engine.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	return g
}

func (g *GinWrapper) WithHealthCheck() *GinWrapper {
	g.engine.GET("/hc", func(c *gin.Context) {
		c.Status(200)
	})
	return g
}

func (g *GinWrapper) WithRateLimiter(opts ...*rateGin.LimitOptions) *GinWrapper {
	r := rate.NewRate()
	limiter := rateGin.Limit(r, opts...)
	g.engine.Use(limiter)
	return g
}

func (g *GinWrapper) WithCors(allowedOrigins []string) *GinWrapper {
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Access-Token"},
		AllowOrigins:     allowedOrigins,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  false,
	}

	g.engine.Use(cors.New(corsConfig))
	return g
}

// Engine returns the proper gin engine
func (g *GinWrapper) Engine() *gin.Engine {
	return g.engine
}
