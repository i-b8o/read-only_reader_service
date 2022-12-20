package service

import (
	"context"
	"read-only_reader_service/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (*entity.Chapter, error)
	GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error)
}

type chapterService struct {
	storage ChapterStorage
}

func NewChapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s *chapterService) Get(ctx context.Context, chapterID uint64) (*entity.Chapter, error) {
	return s.storage.Get(ctx, chapterID)
}

func (s *chapterService) GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error) {
	return s.storage.GetAll(ctx, docID)
}
