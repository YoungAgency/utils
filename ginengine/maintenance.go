package ginengine

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/patrickmn/go-cache"
)

var (
	// MaintenanceModeKey is the key used to store the maintenance mode status in redis
	MaintenanceModeKey = "maintenance"
	redisPool          *redis.Pool
	maintenanceCache   *cache.Cache
)

func init() {
	maintenanceCache = cache.New(1*time.Minute, 1*time.Minute)
}

func SetMaintenanceCacheExpiration(defaultExpiration, cleanUpInterval time.Duration) {
	maintenanceCache = cache.New(defaultExpiration, cleanUpInterval)
}

func SetMaintenanceRedisPool(pool *redis.Pool) {
	redisPool = pool
}

func CheckIfMaintenance(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		maintenance := getStatus(c.Request.Context(), serviceName)

		if maintenance {
			c.AbortWithStatusJSON(
				503, gin.H{"message": "maintenance mode active", "code": "ERR_FEATURE_DISABLED"})
			return
		}
	}
}

func getStatus(ctx context.Context, serviceName string) bool {
	if maintenance, found := maintenanceCache.Get(serviceName); found {
		return maintenance.(bool)
	}

	maintenance := getStatusFromRedis(ctx, serviceName)
	maintenanceCache.Set(serviceName, maintenance, cache.DefaultExpiration)
	return maintenance
}

func getStatusFromRedis(ctx context.Context, serviceName string) bool {
	conn, err := redisPool.GetContext(ctx)
	if err != nil {
		log.Println(err)
		return false
	}
	defer conn.Close()

	reply, err := conn.Do("HGET", "maintenance", serviceName)
	if err != nil {
		log.Println(err)
		return false
	}

	if reply == nil {
		return false
	}

	// if maintenance mode is active, abort request
	var maintenance bool
	if maintenance, err = redis.Bool(reply, nil); err != nil {
		log.Println(err)
		return false
	}
	return maintenance
}
