package reverser

import (
	"context"
	"protobuf/pkg/api"
)

type GRPCServer struct{}

func (s *GRPCServer) Reverse(ctx context.Context, req *api.ReverseRequest) (*api.ReverseResponse, error) {
	return &api.ReverseResponse{Reversed: Reverse(req.GetReversable())}, nil
}
