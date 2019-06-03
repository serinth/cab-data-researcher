package protoServices

import (
	"context"

	"github.com/serinth/cab-data-researcher/cacheService"
	"github.com/serinth/cab-data-researcher/proto"
	log "github.com/sirupsen/logrus"
)

const (
	cacheExpirySecs = int64(15)
)

type cabService struct {
	cache cacheService.CacheService
}

// Function currently stubbed - incomplete
func (service *cabService) GetCabTripsCount(ctx context.Context, id *proto.MedallionId) (*proto.NumberOfTripsResponse, error) {

	service.cache.CacheCabTripCount(context.Background(), "XXX", 1337)
	cnt, err := service.cache.GetTripCount(context.Background(), "XXX")

	log.Infof("count: %v, err: %v", cnt, err)

	return &proto.NumberOfTripsResponse{NumberOfTrips: cnt}, nil
}

func NewCabService(service cacheService.CacheService) proto.CabServer {
	return &cabService{cache: service}
}
