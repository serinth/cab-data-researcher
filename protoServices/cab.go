package protoServices

import (
	"context"
	"encoding/json"
	"time"

	cs "github.com/serinth/cab-data-researcher/cabService"
	"github.com/serinth/cab-data-researcher/cacheService"
	"github.com/serinth/cab-data-researcher/proto"

	log "github.com/sirupsen/logrus"
)

const (
	cacheExpirySecs = int64(15)
	ISO8601Layout = "2006-01-02"
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

	t, _ := time.Parse(ISO8601Layout, "2013-12-01")

	counts, _ := service.cab.GetCabTrips(context.Background(), []string{"D7D598CD99978BD012A87A76A7C891B7", "5455D5FF2BD94D10B304A15D4B7F2735"}, t)

	response, _ := json.Marshal(counts)
	log.Infof("Result: %s", response)

	return &proto.NumberOfTripsResponse{NumberOfTrips: cnt}, nil
}

func NewCabService(service cacheService.CacheService, cabSvc cs.CabService) proto.CabServer {
	return &cabService{cache: service, cab: cabSvc}
}
