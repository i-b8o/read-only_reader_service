package paragraph_controller

import (
	"context"
	"fmt"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type paragraphAdapter interface {
	GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderParagraph, error)
}

type ParagraphGRPCService struct {
	paragraphAdapter paragraphAdapter
	logger           logging.Logger
	pb.UnimplementedParagraphGRPCServer
}

func NewParagraphGRPCService(paragraphAdapter paragraphAdapter, logger logging.Logger) *ParagraphGRPCService {
	return &ParagraphGRPCService{
		paragraphAdapter: paragraphAdapter,
		logger:           logger,
	}
}

func (s *ParagraphGRPCService) GetAll(ctx context.Context, req *pb.GetAllParagraphsByChapterIdRequest) (*pb.GetAllParagraphsByChapterIdResponse, error) {
	id := req.GetID()
	paragraphs, err := s.paragraphAdapter.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(paragraphs) == 0 {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return &pb.GetAllParagraphsByChapterIdResponse{Paragraphs: paragraphs}, nil
}
