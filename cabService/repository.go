package cabService

import (
	"context"
	"time"

	"github.com/serinth/cab-data-researcher/app"
	log "github.com/sirupsen/logrus"
)

// mockgen -source=repository.go -destination=mock_repository.go -package=cabService
type CabRepository interface {
	GetNumberOfTripsByMedallionId(ctx context.Context, ids []string, date time.Time) (int64, error)
}

func NewCabRepository(config *app.Config) CabRepository {
	log.Info("Initializing new Cab Repository instance")
	return &cabMySQL{DB: app.NewDB(config.CabDbConnectionString)}
}
