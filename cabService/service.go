package cabService

import (
	"context"
	"fmt"
	"github.com/serinth/cab-data-researcher/proto"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// mockgen -source=service.go -destination=mock_service_repository.go -package=cabService
type CabService interface {
	GetCabTrips(ctx context.Context, medallionIds []string, date time.Time) ([]*proto.CabTripCount, error)
}

type cabServiceImpl struct {
	repo CabRepository
}

func NewCabService(repository CabRepository) (CabService, error) {
	return &cabServiceImpl{repo: repository}, nil
}

func (service *cabServiceImpl) GetCabTrips(ctx context.Context, medallionIds []string, date time.Time) ([]*proto.CabTripCount, error) {

	results, err := service.repo.GetNumberOfTripsByMedallionIds(context.Background(), medallionIds, date)
	if err != nil {
		log.Errorf("Failed to fetch trips with error %v", err)
		return nil, fmt.Errorf("Failed to get trips for %v on %s", medallionIds, date.String())
	}

	if len(results) == 0 {
		return []*proto.CabTripCount{}, nil
	}

	_, hasCount := results[0]["count"];
	_, hasMedallion := results[0]["medallion"];
	if !hasCount || !hasMedallion {
		return nil, fmt.Errorf("Query did not return with expected results")
	}

	return mapResultsToTripCount(results)
}

func mapResultsToTripCount(queryResults []map[string]string) ([]*proto.CabTripCount, error) {
	var counts []*proto.CabTripCount

	for _, c := range queryResults {
		count, err := strconv.ParseInt(c["count"], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert %s to int64 with error: %v", c["count"], err)
		}
		counts = append(counts, &proto.CabTripCount{Medallion: c["medallion"], Count: count})
	}

	return counts, nil
}
