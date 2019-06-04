package protoServices

import (
	"context"

	cs "github.com/serinth/cab-data-researcher/cabService"
	"github.com/serinth/cab-data-researcher/cacheService"
	"github.com/serinth/cab-data-researcher/proto"

	log "github.com/sirupsen/logrus"
)

const (
	cacheExpirySecs = int64(15)
)

type cabService struct {
	cache cacheService.CacheService
	cab   cs.CabService
}

// Function currently stubbed - incomplete
func (service *cabService) GetCabTripsCount(ctx context.Context, id *proto.MedallionId) (*proto.NumberOfTripsResponse, error) {

	service.cache.CacheCabTripCount(context.Background(), "XXX", 1337)
	cnt, err := service.cache.GetTripCount(context.Background(), "XXX")

	log.Infof("count: %v, err: %v", cnt, err)

	service.cab.GetCabTrips(context.Background())

	return &proto.NumberOfTripsResponse{NumberOfTrips: cnt}, nil
}

func NewCabService(service cacheService.CacheService, cabSvc cs.CabService) proto.CabServer {
	return &cabService{cache: service, cab: cabSvc}
}
