package cabService

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/serinth/cab-data-researcher/proto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCabServiceImpl_GetCabTrips_Returns_Error_On_Failed_Repo_Fetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idsStub := []string{"id1", "id2", "id3"}
	dateStub, _ := time.Parse("2006-01-02", "2013-12-01")

	mockCabRepo := NewMockCabRepository(ctrl)
	mockCabRepo.
		EXPECT().
		GetNumberOfTripsByMedallionIds(gomock.Any(), gomock.Eq(idsStub), gomock.Eq(dateStub)).
		Return(nil, errors.New("error"))

	cabService, _ := NewCabService(mockCabRepo)

	counts, err := cabService.GetCabTrips(context.Background(), idsStub, dateStub)

	assert.Nil(t, counts)
	assert.NotNil(t, err)
}

func TestCabServiceImpl_GetCabTrips_Returns_Empty_On_No_Query_Results(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idsStub := []string{"id1", "id2", "id3"}
	dateStub, _ := time.Parse("2006-01-02", "2013-12-01")
	expected := []*proto.CabTripCount{}

	mockCabRepo := NewMockCabRepository(ctrl)
	mockCabRepo.
		EXPECT().
		GetNumberOfTripsByMedallionIds(gomock.Any(), gomock.Eq(idsStub), gomock.Eq(dateStub)).
		Return([]map[string]string{}, nil)

	cabService, _ := NewCabService(mockCabRepo)

	counts, err := cabService.GetCabTrips(context.Background(), idsStub, dateStub)

	assert.Nil(t, err)
	assert.EqualValues(t, expected, counts)
}

func TestCabServiceImpl_GetCabTrips_Returns_Empty_Error_On_Bad_Query(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idsStub := []string{"id1", "id2", "id3"}
	dateStub, _ := time.Parse("2006-01-02", "2013-12-01")

	mockCabRepo := NewMockCabRepository(ctrl)
	mockCabRepo.
		EXPECT().
		GetNumberOfTripsByMedallionIds(gomock.Any(), gomock.Eq(idsStub), gomock.Eq(dateStub)).
		Return([]map[string]string{{"medallion": "XXX", "errorKey": "error"}}, nil)

	cabService, _ := NewCabService(mockCabRepo)

	counts, err := cabService.GetCabTrips(context.Background(), idsStub, dateStub)

	assert.NotNil(t, err)
	assert.Nil(t, counts)
}

func TestCabServiceImpl_GetCabTrips(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idsStub := []string{"id1", "id2", "id3"}
	dateStub, _ := time.Parse("2006-01-02", "2013-12-01")
	expected := []*proto.CabTripCount{
		{Medallion: "XXX", Count: 42},
	}

	mockCabRepo := NewMockCabRepository(ctrl)
	mockCabRepo.
		EXPECT().
		GetNumberOfTripsByMedallionIds(gomock.Any(), gomock.Eq(idsStub), gomock.Eq(dateStub)).
		Return([]map[string]string{{"medallion": "XXX", "count": "42"}}, nil)

	cabService, _ := NewCabService(mockCabRepo)

	counts, err := cabService.GetCabTrips(context.Background(), idsStub, dateStub)

	assert.Nil(t, err)
	assert.EqualValues(t, expected, counts)
}


func TestCabServiceImpl_MapResultsToTripCount(t *testing.T) {
	var mapperTests = []struct {
		name         string
		queryResults []map[string]string
		expected     []*proto.CabTripCount
	}{
		{
			"single good case",
			[]map[string]string{{"medallion": "XXX", "count": "42"}},
			[]*proto.CabTripCount{{Medallion: "XXX", Count: 42}},
		},
		{
			"multiple good case",
			[]map[string]string{
				{"medallion": "XXX", "count": "42"},
				{"medallion": "YYY", "count": "1337"},
				{"medallion": "ZZZ", "count": "42"},
			},
			[]*proto.CabTripCount{
				{Medallion: "XXX", Count: 42},
				{Medallion: "YYY", Count: 1337},
				{Medallion: "ZZZ", Count: 42},
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

func TestCabServiceImpl_MapResultsToTripCount_Should_Error(t *testing.T) {
	test := []map[string]string{{"medallion": "XXX", "count": "error"}}

	results, err := mapResultsToTripCount(test)
	assert.NotNil(t, err)
	assert.Nil(t, results)
}
