package adder

import (
	"context"

	"grpc-go-sum/pkg/api"
)

type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
