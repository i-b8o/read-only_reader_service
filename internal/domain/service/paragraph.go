package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
}

type paragraphService struct {
	storage ParagraphStorage
	logger  logging.Logger
}

func NewParagraphService(storage ParagraphStorage, logger logging.Logger) *paragraphService {
	return &paragraphService{storage: storage, logger: logger}
}

func (s paragraphService) GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error) {
	paragraphs, err := s.storage.GetAll(ctx, chapterID)
	if err != nil {
		s.logger.Errorf("could not get paragraphs for %d, %v", chapterID, err)
		return nil, err
	}
	return paragraphs, nil
}
