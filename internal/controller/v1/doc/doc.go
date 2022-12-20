package doc_controller

import (
	"context"
	"fmt"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DocUsecase interface {
	Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error)
}

type DocGRPCService struct {
	docUsecase DocUsecase
	pb.UnimplementedDocGRPCServer
}

func NewDocGRPCService(docUsecase DocUsecase) *DocGRPCService {
	return &DocGRPCService{
		docUsecase: docUsecase,
	}
}

func (s *DocGRPCService) GetOne(ctx context.Context, req *pb.GetOneDocRequest) (*pb.GetOneDocResponse, error) {
	id := req.GetID()
	resp, err := s.docUsecase.Get(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return resp, nil
}
