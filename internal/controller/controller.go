package controller

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

type ChapterUsecase interface {
	Get(ctx context.Context, chapterID uint64) (*pb.GetOneChapterResponse, error)
	GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReaderChapter, error)
}

// type ParagraphUsecase interface {
// 	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
// }

type ReadOnlyRegulationGRPCService struct {
	regulationUsecase RegulationUsecase
	chapterUsecase    ChapterUsecase

	pb.UnimplementedReaderGRPCServer
}

func NewReaderGRPCService(regulationUsecase RegulationUsecase, chapterUsecase ChapterUsecase) *ReadOnlyRegulationGRPCService {
	return &ReadOnlyRegulationGRPCService{
		regulationUsecase: regulationUsecase,
		chapterUsecase:    chapterUsecase,
	}
}

func (s *ReadOnlyRegulationGRPCService) GetOneRegulation(ctx context.Context, req *pb.GetOneRegulationRequest) (*pb.GetOneRegulationResponse, error) {
	id := req.GetID()
	resp, err := s.regulationUsecase.Get(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return resp, nil
}

func (s *ReadOnlyRegulationGRPCService) GetOneChapter(ctx context.Context, req *pb.GetOneChapterRequest) (*pb.GetOneChapterResponse, error) {
	id := req.GetID()
	chapter, err := s.chapterUsecase.Get(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}

	return chapter, nil
}

func (s *ReadOnlyRegulationGRPCService) GetAllChaptersByRegulationId(ctx context.Context, req *pb.GetAllChaptersByRegulationIdRequest) (*pb.GetAllChaptersByRegulationIdResponse, error) {
	id := req.GetID()
	chapters, err := s.chapterUsecase.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(chapters) == 0 {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return &pb.GetAllChaptersByRegulationIdResponse{Chapters: chapters}, nil
}
