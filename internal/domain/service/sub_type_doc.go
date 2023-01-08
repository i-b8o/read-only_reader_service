package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type SubTypeDocStorage interface {
	GetAll(ctx context.Context, subtypeID uint64) ([]*pb.Doc, error)
}
type subTypeDocService struct {
	storage SubTypeDocStorage
	logger  logging.Logger
}

func NewSubTypeDocService(storage SubTypeDocStorage, logger logging.Logger) *subTypeDocService {
	return &subTypeDocService{storage: storage, logger: logger}
}

func (s *subTypeDocService) GetAll(ctx context.Context, subtypeID uint64) (*pb.GetDocsResponse, error) {
	docs, err := s.storage.GetAll(ctx, subtypeID)
	if err != nil {
		s.logger.Errorf("could not get docs for %d, %v", subtypeID, err)
		return nil, err
	}
	return &pb.GetDocsResponse{Docs: docs}, nil
}
