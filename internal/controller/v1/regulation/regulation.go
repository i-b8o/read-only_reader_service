package regulation_controller

import (
	"context"
	"fmt"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegulationUsecase interface {
	Get(ctx context.Context, regulationID uint64) (*pb.GetOneRegulationResponse, error)
}

type RegulationGRPCService struct {
	regulationUsecase RegulationUsecase
	pb.UnimplementedRegulationGRPCServer
}

func NewRegulationGRPCService(regulationUsecase RegulationUsecase) *RegulationGRPCService {
	return &RegulationGRPCService{
		regulationUsecase: regulationUsecase,
	}
}

func (s *RegulationGRPCService) GetOne(ctx context.Context, req *pb.GetOneRegulationRequest) (*pb.GetOneRegulationResponse, error) {
	id := req.GetID()
	resp, err := s.regulationUsecase.Get(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return resp, nil
}
