package cabService

import (
	"context"

	"github.com/serinth/cab-data-researcher/app"
	log "github.com/sirupsen/logrus"
)

// go:generate moq -out mock_ecomm_repository.test.go . EcommRepository
type CabRepository interface {
	GetNumberOfTripsByMedallionId(ctx context.Context, id string) (int64, error)
}

func NewEcommCloudRepository(config *app.Config) CabRepository {
	log.Info("Initializing new Cab Repository instance")
	return &cabMySQL{DB: app.NewDB(config.CabDbConnectionString)}
}
