package cacheService

import (
	"context"
)

type CacheService interface {
	CacheCabTripCount(ctx context.Context, id string, count int64) error
	GetTripCount(ctx context.Context, id string) (int64, error)
}

type cacheServiceImpl struct {
	repo CacheRepository
}

func NewCacheService(repository CacheRepository) (CacheService, error) {
	return &cacheServiceImpl{repo: repository}, nil
}

func (service *cacheServiceImpl) CacheCabTripCount(ctx context.Context, id string, count int64) error {
	return nil
}

func (service *cacheServiceImpl) GetTripCount(ctx context.Context, id string) (int64, error) {
	return 0, nil
}
