package cabService

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	CabSchema     = "ny_cab_data"
	TripsTable    = "cab_trip_data"
)

type cabMySQL struct {
	DB xorm.EngineInterface
}


func (repo *cabMySQL) GetNumberOfTripsByMedallionIds(ctx context.Context, ids []string, date time.Time) ([]map[string]string, error) {
	return repo.DB.
		Table(fmt.Sprintf("%s.%s", CabSchema, TripsTable)).
		Select("medallion, COUNT(medallion) as count").
		In("medallion", ids).
		Where("DATE(pickup_datetime) = ?", date).
		GroupBy("medallion").
		QueryString()
}
