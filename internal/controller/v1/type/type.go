package type_controller

import (
	"context"
	"fmt"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TypeAdapter interface {
	GetAll(ctx context.Context) ([]*pb.TypeResponse, error)
}

type TypeGRPCService struct {
	typeAdapter TypeAdapter
	logger      logging.Logger
	pb.UnimplementedTypeGRPCServer
}

func NewTypeGRPCService(typeAdapter TypeAdapter, logger logging.Logger) *TypeGRPCService {
	return &TypeGRPCService{
		typeAdapter: typeAdapter,
		logger:      logger,
	}
}

func (s *TypeGRPCService) GetAll(ctx context.Context, req *pb.Empty) (*pb.GetAllTypesResponse, error) {
	resp, err := s.typeAdapter.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("could not get types %v", err)
		err := status.Errorf(codes.NotFound, fmt.Sprintf("could not get types: %v", err))
		return nil, err
	}
	return &pb.GetAllTypesResponse{Types: resp}, nil
}
