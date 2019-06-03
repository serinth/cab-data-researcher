package cacheService

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	expiry = int64(1337)
)

func TestCacheService_CacheCabTripCount_ReturnsError_When_RedisClient_Errors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedis := NewMockCacheRepository(ctrl)
	mockRedis.
		EXPECT().
		SaveExpiringKeyValue(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(expiry)).
		Return(errors.New("error")).
		Times(1)

	cacheService, _ := NewCacheService(mockRedis, expiry)

	err := cacheService.CacheCabTripCount(context.Background(), "key", 42)
	assert.NotNil(t, err)
}

func TestCacheService_CacheCabTripCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRedis := NewMockCacheRepository(ctrl)
	mockRedis.
		EXPECT().
		SaveExpiringKeyValue(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(expiry)).
		Return(nil)

	cacheService, _ := NewCacheService(mockRedis, expiry)

	err := cacheService.CacheCabTripCount(context.Background(), "key", 42)
	assert.Nil(t, err)
}
