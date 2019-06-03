package cabService

import (
	"context"
)

type CabService interface {
	GetCabTrips(ctx context.Context) (int64, error)
}

type cabServiceImpl struct {
	repo CabRepository
}

func NewCabService(repository CabRepository) (CabService, error) {
	return &cabServiceImpl{ repo: repository }, nil
}

func (service *cabServiceImpl) GetCabTrips(ctx context.Context) (int64, error) {
	return 0, nil
}