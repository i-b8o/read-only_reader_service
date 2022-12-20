package paragraph_controller

import (
	"context"
	"fmt"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type paragraphUsecase interface {
	GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderParagraph, error)
}

type ParagraphGRPCService struct {
	paragraphUsecase paragraphUsecase
	pb.UnimplementedParagraphGRPCServer
}

func NewParagraphGRPCService(chapterUsecase paragraphUsecase) *ParagraphGRPCService {
	return &ParagraphGRPCService{
		paragraphUsecase: chapterUsecase,
	}
}

func (s *ParagraphGRPCService) GetAll(ctx context.Context, req *pb.GetAllParagraphsByChapterIdRequest) (*pb.GetAllParagraphsByChapterIdResponse, error) {
	id := req.GetID()
	paragraphs, err := s.paragraphUsecase.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(paragraphs) == 0 {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return &pb.GetAllParagraphsByChapterIdResponse{Paragraphs: paragraphs}, nil
}
