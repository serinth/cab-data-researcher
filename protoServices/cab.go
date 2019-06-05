package protoServices

import (
	"context"
	"github.com/serinth/cab-data-researcher/validators"
	"time"

	cs "github.com/serinth/cab-data-researcher/cabService"
	"github.com/serinth/cab-data-researcher/cacheService"
	"github.com/serinth/cab-data-researcher/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
)

const (
	ISO8601Layout   = "2006-01-02"
)

type cabService struct {
	cache cacheService.CacheService
	cab   cs.CabService
}

// Function currently stubbed - incomplete
func (service *cabService) GetCabTripsCount(ctx context.Context, request *proto.CabTripsRequest) (*proto.CabTripsResponse, error) {

	if len(request.Medallions) == 0 || len(request.Date) == 0 {
		return 	nil, status.Error(codes.InvalidArgument, "Bad Request: Missing arguments")
	}

	var ids []string

	for _, i := range request.Medallions {
		ids = append(ids, i.Id)
	}

	if !validators.ContainsOnlyAlphanumeric(ids) {
		return nil, status.Error(codes.InvalidArgument, "Bad Request, Medallion Ids not in the correct format")
	}

	t, err := time.Parse(ISO8601Layout, request.Date)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s is not in the format YYYY-mm-dd", request.Date)
	}


	var counts []*proto.CabTripCount
	var cacheMiss []string

	if request.SkipCache == true {
		counts, err = service.cab.GetCabTrips(ctx, ids, t)
		if err != nil {
			log.Errorf("Failed to fetch cab trips with error: %v", err)
			return nil, status.Error(codes.Internal, "Internal Error")
		}
		service.cacheCounts(counts)
	} else {
		for _, id := range ids {
			cnt, err := service.cache.GetTripCount(ctx, id)
			if err == nil {
				counts = append(counts, &proto.CabTripCount{Medallion: id, Count: cnt})
			} else {
				log.Infof("Cache Miss / Error for id: %s with error: %v", id, err)
				cacheMiss = append(cacheMiss, id)
			}
		}
	}

	if len(cacheMiss) > 0 {
		missed, err := service.cab.GetCabTrips(ctx, cacheMiss, t)
		if err != nil {
			log.Errorf("Failed to fetch cab trips for cache misses: %v with error %v", cacheMiss, err)
			return nil, status.Error(codes.Internal, "Internal Error")
		}
		service.cacheCounts(missed)

		for _, i := range missed {
			counts = append(counts, i)
		}
	}

	return &proto.CabTripsResponse{Results: counts}, nil
}

func (service *cabService) cacheCounts(counts []*proto.CabTripCount) {
	ctx := context.Background()

	for _, i := range counts {
		err := service.cache.CacheCabTripCount(ctx, i.Medallion, i.Count)
		if err != nil {
			log.Warnf("Failed to cache %s with value %d", i.Medallion, i.Count)
		}
	}
}

func NewCabService(service cacheService.CacheService, cabSvc cs.CabService) proto.CabServer {
	return &cabService{cache: service, cab: cabSvc}
}
