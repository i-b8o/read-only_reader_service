package service

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type RegulationStorage interface {
	Get(ctx context.Context, regulationID uint64) (*pb.GetOneRegulationResponse, error)
}

type regulationService struct {
	storage RegulationStorage
}

func NewRegulationService(storage RegulationStorage) *regulationService {
	return &regulationService{storage: storage}
}

func (s *regulationService) Get(ctx context.Context, regulationID uint64) (*pb.GetOneRegulationResponse, error) {
	return s.storage.Get(ctx, regulationID)
}
