package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (*pb.ReaderChapter, error)
	GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error)
}

type chapterService struct {
	storage ChapterStorage
	logger  logging.Logger
}

func NewChapterService(storage ChapterStorage, logger logging.Logger) *chapterService {
	return &chapterService{storage: storage, logger: logger}
}

func (s *chapterService) Get(ctx context.Context, chapterID uint64) (*pb.ReaderChapter, error) {
	chapter, err := s.storage.Get(ctx, chapterID)
	if err != nil {
		s.logger.Errorf("could not get chapter for %d, %v", chapterID, err)
		return nil, err
	}
	return chapter, nil
}

func (s *chapterService) GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error) {
	chapters, err := s.storage.GetAll(ctx, docID)
	if err != nil {
		s.logger.Errorf("could not get chapters for %d, %v", docID, err)
		return nil, err
	}
	return chapters, nil
}
