package chapter_controller

import (
	"context"
	"fmt"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type ChapterUsecase interface {
// 	Get(ctx context.Context, chapterID uint64) (*pb.GetOneChapterResponse, error)
// 	GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error)
// }

type ChapterAdapter interface {
	Get(ctx context.Context, chapterID uint64) (*pb.ReaderChapter, error)
	GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error)
}

type ParagraphAdapter interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
}

type ChapterGRPCService struct {
	chapterAdapter   ChapterAdapter
	paragraphAdapter ParagraphAdapter
	logger           logging.Logger
	pb.UnimplementedChapterGRPCServer
}

func NewChapterGRPCService(chapterAdapter ChapterAdapter, paragraphAdapter ParagraphAdapter, logger logging.Logger) *ChapterGRPCService {
	return &ChapterGRPCService{
		chapterAdapter:   chapterAdapter,
		paragraphAdapter: paragraphAdapter,
		logger:           logger,
	}
}

func (s *ChapterGRPCService) GetOne(ctx context.Context, req *pb.GetOneChapterRequest) (*pb.GetOneChapterResponse, error) {
	id := req.GetID()
	chapterInfo, err := s.chapterAdapter.Get(ctx, id)
	if err != nil {
		s.logger.Errorf("could not get chapter for %d, %v", id, err)
		return nil, err
	}
	paragraphs, err := s.paragraphAdapter.GetAll(ctx, id)
	if err != nil {
		s.logger.Errorf("could not get paragraphs for %d, %v", id, err)
		return nil, err
	}
	return &pb.GetOneChapterResponse{ID: id, Name: chapterInfo.Name, Num: chapterInfo.Num, DocID: chapterInfo.DocID, OrderNum: chapterInfo.OrderNum, Paragraphs: paragraphs}, nil
}

func (s *ChapterGRPCService) GetAll(ctx context.Context, req *pb.GetAllChaptersByDocIdRequest) (*pb.GetAllChaptersByDocIdResponse, error) {
	id := req.GetID()
	chapters, err := s.chapterAdapter.GetAll(ctx, id)
	if err != nil {
		s.logger.Errorf("could not get chapters for %d, %v", id, err)
		return nil, err
	}
	if len(chapters) == 0 {
		err := status.Errorf(codes.NotFound, fmt.Sprintf("could not get chapters for %d, %v", id, err))
		return nil, err
	}
	return &pb.GetAllChaptersByDocIdResponse{Chapters: chapters}, nil
}
