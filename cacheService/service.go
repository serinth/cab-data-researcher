package cacheService

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type CacheService interface {
	CacheCabTripCount(ctx context.Context, id string, count int64) error
	GetTripCount(ctx context.Context, id string) (int64, error)
}

type cacheServiceImpl struct {
	repo        CacheRepository
	cacheExpiry int64
}

func NewCacheService(repository CacheRepository, cacheExpiry int64) (CacheService, error) {
	return &cacheServiceImpl{repo: repository, cacheExpiry: cacheExpiry}, nil
}

func (service *cacheServiceImpl) CacheCabTripCount(ctx context.Context, id string, count int64) error {
	// can repeatedly cache without error, will update the expiry
	err := service.repo.SaveExpiringKeyValue(ctx, id, count, service.cacheExpiry)

	if err != nil {
		log.Errorf("Failed to cache count: %d for key: %s with error: %v", count, id, err)
		return err
	}

	return nil
}

func (service *cacheServiceImpl) GetTripCount(ctx context.Context, id string) (int64, error) {

	count, err := service.repo.GetKey(ctx, id)

	if err != nil {
		log.Infof("%v", err)
		return 0, err
	}

	return count, nil
}
