package chapter_controller

import (
	"context"
	"fmt"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChapterUsecase interface {
	Get(ctx context.Context, chapterID uint64) (*pb.GetOneChapterResponse, error)
	GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error)
}

type ChapterGRPCService struct {
	chapterUsecase ChapterUsecase
	pb.UnimplementedChapterGRPCServer
}

func NewChapterGRPCService(chapterUsecase ChapterUsecase) *ChapterGRPCService {
	return &ChapterGRPCService{
		chapterUsecase: chapterUsecase,
	}
}

func (s *ChapterGRPCService) GetOne(ctx context.Context, req *pb.GetOneChapterRequest) (*pb.GetOneChapterResponse, error) {
	id := req.GetID()
	chapter, err := s.chapterUsecase.Get(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}

	return chapter, nil
}

func (s *ChapterGRPCService) GetAll(ctx context.Context, req *pb.GetAllChaptersByDocIdRequest) (*pb.GetAllChaptersByDocIdResponse, error) {
	id := req.GetID()
	chapters, err := s.chapterUsecase.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(chapters) == 0 {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("id was not found: %d", id))
		return nil, err
	}
	return &pb.GetAllChaptersByDocIdResponse{Chapters: chapters}, nil
}
