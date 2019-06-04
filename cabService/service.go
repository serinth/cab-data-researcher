package cabService

import (
	"context"
	"time"
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
	t, _ := time.Parse(ISO8601Layout, "2013-12-01")
	service.repo.GetNumberOfTripsByMedallionId(context.Background(), []string{"D7D598CD99978BD012A87A76A7C891B7", "5455D5FF2BD94D10B304A15D4B7F2735"}, t)
	return 0, nil
}