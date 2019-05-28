package utils

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/hc"},
		Formatter: func(param gin.LogFormatterParams) string {
			if param.Method == "OPTIONS" {
				return ""
			}
			return fmt.Sprintf("[GIN] %3d | %13v | %15s | %s %-7s\n%s",
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				param.Path,
				param.ErrorMessage,
			)
		},
	})
}

func GinCors(allowedOrigins ...string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowOrigins:     allowedOrigins,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  false,
	})
}
