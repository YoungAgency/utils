package utils

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GinCors(allowedOrigins ...string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Access-Token"},
		AllowOrigins:     allowedOrigins,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  false,
	})
}
