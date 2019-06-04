package cabService

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

const (
	CabSchema     = "ny_cab_data"
	TripsTable    = "cab_trip_data"
	ISO8601Layout = "2006-01-02"
)

type cabMySQL struct {
	DB xorm.EngineInterface
}


func (repo *cabMySQL) GetNumberOfTripsByMedallionId(ctx context.Context, ids []string, date time.Time) (int64, error) {

	/*
	select g.medallion, sum(g.cnt) from (
	select medallion, count(medallion) as cnt, DATE(pickup_datetime) as dt
	from cab_trip_data
	where medallion in (
		'D7D598CD99978BD012A87A76A7C891B7',
		'5455D5FF2BD94D10B304A15D4B7F2735',
		'FFFECF75AB6CC4FF9E8A8B633AB81C26',
		'jibberish')
	group by medallion, dt
) g
group by g.medallion
	 */

/*
	select medallion, count(medallion) as cnt
		from cab_trip_data
		where medallion in (
			'D7D598CD99978BD012A87A76A7C891B7', -- see 3
		'5455D5FF2BD94D10B304A15D4B7F2735', -- see 2
		'801C69A08B51470871A8110F8B0505EE', -- see 1
		'5455D5FF2BD94D10B304A15D4B7F2735')
		and DATE(pickup_datetime) = '2013-12-01'
	group by medallion


		*/


	results, err := repo.DB.
		Table(fmt.Sprintf("%s.%s", CabSchema, TripsTable)).
		Select("medallion, COUNT(medallion) as count").
		In("medallion", ids).
		Where("DATE(pickup_datetime) = ?", date.Format(ISO8601Layout)).
		GroupBy("medallion").
		QueryString()


	for _, i := range results {
		log.Infof("QUERY RESULT: %v, %v", i["count"], i["medallion"])
	}



	if err != nil {
		log.Errorf("Failed to get trip counts with error: %v", err)
		return 0, err
	}

	//return count, nil

	return 0, nil
}

//func formatSelectIn(values []string) *betweenDateTuple {
//
//	startDate := date.Format(ISO8601Layout)
//	endDate := date.AddDate(0, 0, 1).Format(ISO8601Layout)
//
//	return &betweenDateTuple{
//		start: startDate,
//		end:   endDate,
//	}
//}
