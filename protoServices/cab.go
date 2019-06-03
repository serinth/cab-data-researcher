package protoServices

import (
	"context"

	"github.com/serinth/cab-data-researcher/proto"
)

type cabService struct {}

func (service *cabService) GetCabTripsCount(ctx context.Context, id *proto.MedallionId) (*proto.NumberOfTripsResponse, error){
	return &proto.NumberOfTripsResponse{NumberOfTrips: 1}, nil
}

func NewCabService() proto.CabServer {
	return new(cabService)
}