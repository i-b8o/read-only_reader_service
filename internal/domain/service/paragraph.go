package service

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
}

type paragraphService struct {
	storage ParagraphStorage
}

func NewParagraphService(storage ParagraphStorage) *paragraphService {
	return &paragraphService{storage: storage}
}

func (s paragraphService) GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error) {
	return s.storage.GetAll(ctx, chapterID)
}
