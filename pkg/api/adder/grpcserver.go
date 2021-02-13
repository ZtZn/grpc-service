package adder

import (
	"context"

	"github.com/ztzn/grpc_services/pkg/api"
)

type GRPCServer struct {
}

// Add..
func (s *GRPCServer) Add(ctx context.Context, in *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: in.GetX() + in.GetY()}, nil
}
