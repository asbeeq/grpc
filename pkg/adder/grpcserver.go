package adder

import (
	"context"

	"github.com/asbeeq/grpc/pkg/api"
)

type GRPCServer struct{}

func (g *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func (g *GRPCServer) Sub(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() - req.GetY()}, nil
}
