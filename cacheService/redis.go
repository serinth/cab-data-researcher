package cacheService

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type redisClient struct {
	client *redis.Client
}

func (r *redisClient) SaveExpiringKeyValue(ctx context.Context, key string, value int64, expireSeconds int64) error {
	expiryTime := time.Duration(expireSeconds) * time.Second

	err := r.client.Set(key, value, expiryTime).Err()
	if err != nil {
		log.Errorf("Failed to cache count for %s with error: %v", key, err)
		return err
	}

	return nil
}

func (r *redisClient) GetKey(ctx context.Context, key string) (int64, error) {

	val, err := r.client.Get(key).Result()

	if err != nil {
		log.Errorf("Failed to get from cache for key: %s with error: %v", key, err)
		return 0, err
	}

	if err == redis.Nil {
		return 0, nil
	}

	count, err := strconv.ParseInt(val, 10, 64)

	if err != nil {
		log.Errorf("Failed to convert %s to int from redis with error: %v", val, err)
		return 0, err
	}

	return count, nil
}
