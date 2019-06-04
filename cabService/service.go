package cabService

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/serinth/cab-data-researcher/cabService/models"
	log "github.com/sirupsen/logrus"
)

type CabService interface {
	GetCabTrips(ctx context.Context, medallionIds []string, date time.Time) ([]*models.TripCount, error)
}

type cabServiceImpl struct {
	repo CabRepository
}

func NewCabService(repository CabRepository) (CabService, error) {
	return &cabServiceImpl{repo: repository}, nil
}

func (service *cabServiceImpl) GetCabTrips(ctx context.Context, medallionIds []string, date time.Time) ([]*models.TripCount, error) {
	//t, err := time.Parse(ISO8601Layout, date)

	results, err := service.repo.GetNumberOfTripsByMedallionIds(context.Background(), medallionIds, date)
	if err != nil {
		log.Errorf("Failed to fetch trips with error %v", err)
		return nil, fmt.Errorf("Failed to get trips for %v on %s", medallionIds, date.String())
	}

	if len(results) == 0 {
		return []*models.TripCount{}, nil
	}

	_, hasCount := results[0]["count"];
	_, hasMedallion := results[0]["medallion"];

	if !hasCount || !hasMedallion {
		return nil, fmt.Errorf("Query did not return with expected results")
	}

	return mapResultsToTripCount(results)
}

func mapResultsToTripCount(queryResults []map[string]string) ([]*models.TripCount, error) {
	var counts []*models.TripCount

	for _, c := range queryResults {
		count, err := strconv.Atoi(c["count"])
		if err != nil {
			return nil, fmt.Errorf("Failed to convert %s to int with error: %v", c["count"], err)
		}
		counts = append(counts, &models.TripCount{Medallion: c["medallion"], Count: count})
	}

	return counts, nil
}
