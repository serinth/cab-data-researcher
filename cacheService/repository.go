package cacheService

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/serinth/cab-data-researcher/app"
	log "github.com/sirupsen/logrus"
)

// mockgen -source=repository.go -destination=mock_repository.go -package=cacheService
type CacheRepository interface {
	SaveExpiringKeyValue(ctx context.Context, key string, value int64, expireSeconds int64) error
	GetKey(ctx context.Context, key string) (int64, error)
}

func NewCacheRepository(config *app.Config) CacheRepository {
	log.Info("Initializing new Redis Cache")

	c := redis.NewClient(&redis.Options{
		Addr:     config.RedisURI,
		Password: config.RedisPassword,
		DB:       config.RedisDb,
	})

	_, err := c.Ping().Result()

	if err != nil {
		return nil
	}

	return &redisClient{client: c}
}
