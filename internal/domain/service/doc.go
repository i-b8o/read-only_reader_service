package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type DocStorage interface {
	Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error)
}

type docService struct {
	storage DocStorage
	logger  logging.Logger
}

func NewDocService(storage DocStorage, logger logging.Logger) *docService {
	return &docService{storage: storage, logger: logger}
}

func (s *docService) Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error) {
	doc, err := s.storage.Get(ctx, docID)
	if err != nil {
		s.logger.Errorf("could not get doc for %d, %v", docID, err)
		return nil, err
	}
	return doc, nil
}
