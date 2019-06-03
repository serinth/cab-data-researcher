package cabService

import (
	"context"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//log "github.com/sirupsen/logrus"
)

const (
	CabSchema      = "ny_cab_data"
	SalesOrderTable = "cab_trip_data"
)

type CabMySQL struct {
	DB xorm.EngineInterface
}

func (repo *CabMySQL) GetNumberOfTripsByMedallionId(ctx context.Context, id string) (int64, error) {
	return 0, nil
}