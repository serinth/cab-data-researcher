package cacheService

import (
	"context"

	"github.com/go-redis/redis"
)

type redisClient struct {
	client *redis.Client
}

func (redis *redisClient) SaveKeyValue(ctx context.Context, key string, value int64) error {
	return nil
}

func (redis *redisClient) GetKey(ctx context.Context, key string) (int64, error) {
	return 0, nil
}
