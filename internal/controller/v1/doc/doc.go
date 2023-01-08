package doc_controller

import (
	"context"
	"fmt"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DocAdapter interface {
	Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error)
}

type DocGRPCService struct {
	docAdapter DocAdapter
	logger     logging.Logger
	pb.UnimplementedDocGRPCServer
}

func NewDocGRPCService(docAdapter DocAdapter, logger logging.Logger) *DocGRPCService {
	return &DocGRPCService{
		docAdapter: docAdapter,
		logger:     logger,
	}
}

func (s *DocGRPCService) GetOne(ctx context.Context, req *pb.GetOneDocRequest) (*pb.GetOneDocResponse, error) {
	id := req.GetID()
	resp, err := s.docAdapter.Get(ctx, id)
	if err != nil {
		s.logger.Errorf("could not get doc for %d, %v", id, err)
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return resp, nil
}
