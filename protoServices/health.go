package protoServices

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/serinth/cab-data-researcher/cab-data-researcher/proto"
)

type healthService struct {}

func (service *healthService) HealthCheck(context.Context, *empty.Empty) (*proto.HealthStatus, error) {
	return &proto.HealthStatus{Status: "OK"}, nil
}

// NewHealthService - Returns new implementation of health service
func NewHealthService() proto.HealthServer {
	return new(healthService)
}