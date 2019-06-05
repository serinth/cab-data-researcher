package protoServices

import (
	"context"
	"github.com/golang/mock/gomock"
	cs "github.com/serinth/cab-data-researcher/cabService"
	"github.com/serinth/cab-data-researcher/cacheService"
	"github.com/serinth/cab-data-researcher/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCabService_GetCabTripsCount_ReturnsError_On_Empty_MedallionIds(t *testing.T) {
	mockCacheService, mockCabService := makeMocks(t)
	ctx := context.Background()

	request := &proto.CabTripsRequest{}

	service := NewCabService(mockCacheService, mockCabService)

	res, err := service.GetCabTripsCount(ctx, request)

	assert.Nil(t, res)
	assert.EqualValues(t, status.Error(codes.InvalidArgument, "Bad Request: Missing arguments"), err)
}

func TestCabService_GetCabTripsCount_ReturnsError_On_Bad_Ids(t *testing.T) {
	mockCacheService, mockCabService := makeMocks(t)
	ctx := context.Background()

	request := &proto.CabTripsRequest{
		Medallions: []*proto.MedallionId{{"$$"}},
		Date:       "2013-12-01",
		SkipCache:  true,
	}

	service := NewCabService(mockCacheService, mockCabService)

	res, err := service.GetCabTripsCount(ctx, request)

	assert.Nil(t, res)
	assert.EqualValues(t, status.Error(codes.InvalidArgument, "Bad Request, Medallion Ids not in the correct format"), err)
}

func TestCabService_GetCabTripsCount_ReturnsError_On_Bad_Date(t *testing.T) {
	mockCacheService, mockCabService := makeMocks(t)
	ctx := context.Background()

	request := &proto.CabTripsRequest{
		Medallions: []*proto.MedallionId{{"ABCD123"}},
		Date:       "error",
		SkipCache:  false,
	}

	service := NewCabService(mockCacheService, mockCabService)

	res, err := service.GetCabTripsCount(ctx, request)

	assert.Nil(t, res)
	assert.EqualValues(t, status.Errorf(codes.InvalidArgument, "%s is not in the format YYYY-mm-dd", request.Date), err)
}

func TestCabService_GetCabTripsCount_SkipCache_True(t *testing.T) {
	mockCacheService, mockCabService := makeMocks(t)
	ctx := context.Background()
	tripCountStub := []*proto.CabTripCount{{Medallion: "XXX", Count: 42}}
	expected := &proto.CabTripsResponse{Results: tripCountStub}

	request := &proto.CabTripsRequest{
		Medallions: []*proto.MedallionId{{"XXX"}},
		Date:       "2013-12-01",
		SkipCache:  true,
	}

	mockCacheService.
		EXPECT().
		CacheCabTripCount(gomock.Any(), gomock.Eq("XXX"), gomock.Eq(int64(42))).
		Times(1)

	mockCacheService.
		EXPECT().
		GetTripCount(gomock.Any(), gomock.Any()).
		Times(0)

	mockCabService.
		EXPECT().
		GetCabTrips(gomock.Any(), gomock.Eq([]string{"XXX"}), gomock.Any()).
		Return(tripCountStub, nil)

	service := NewCabService(mockCacheService, mockCabService)

	res, err := service.GetCabTripsCount(ctx, request)

	assert.Nil(t, err)
	assert.EqualValues(t, expected, res)
}

// TODO more tests like the above, just modify the mocks and their expectations

func makeMocks(t *testing.T) (*cacheService.MockCacheService, *cs.MockCabService) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCacheService := cacheService.NewMockCacheService(ctrl)
	mockCabService := cs.NewMockCabService(ctrl)

	return mockCacheService, mockCabService
}
