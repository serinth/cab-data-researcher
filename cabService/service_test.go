package cabService

import (
	"github.com/serinth/cab-data-researcher/cabService/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapResultsToTripCount(t *testing.T) {
	var mapperTests = []struct {
		name         string
		queryResults []map[string]string
		expected     []*models.TripCount
	}{
		{
			"single good case",
			[]map[string]string{{"medallion": "XXX", "count": "42"}},
			[]*models.TripCount{&models.TripCount{"XXX", 42}},
		},
		{
			"multiple good case",
			[]map[string]string{
				{"medallion": "XXX", "count": "42"},
				{"medallion": "YYY", "count": "1337"},
				{"medallion": "ZZZ", "count": "42"},
			},
			[]*models.TripCount{
				&models.TripCount{"XXX", 42},
				&models.TripCount{"YYY", 1337},
				&models.TripCount{"ZZZ", 42},
			},
		},
	}

	for _, tt := range mapperTests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := mapResultsToTripCount(tt.queryResults)
			assert.EqualValues(t, tt.expected, results)
			assert.Nil(t, err)
		})
	}
}

func TestMapResultsToTripCount_Should_Error(t *testing.T) {
	test := []map[string]string{{"medallion": "XXX", "count": "error"}}

	results, err := mapResultsToTripCount(test)
	assert.NotNil(t, err)
	assert.Nil(t, results)
}