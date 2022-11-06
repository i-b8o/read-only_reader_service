package service

import (
	"context"
	"regulations_service/internal/domain/entity"
)

type RegulationStorage interface {
	CreateOne(ctx context.Context, regulation entity.Regulation) (uint64, error)
}

type regulationService struct {
	storage RegulationStorage
}

func NewRegulationService(storage RegulationStorage) *regulationService {
	return &regulationService{storage: storage}
}

func (s *regulationService) CreateOne(ctx context.Context, regulation entity.Regulation) (uint64, error) {
	return s.storage.CreateOne(ctx, regulation)
}

// func (s *regulationService) GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error) {
// 	return s.storage.GetOne(ctx, regulationID)
// }

// func (s *regulationService) GetAll(ctx context.Context) ([]entity.Regulation, error) {
// 	return s.storage.GetAll(ctx)
// }

// func (s *regulationService) DeleteRegulation(ctx context.Context, regulationID uint64) error {
// 	return s.storage.DeleteRegulation(ctx, regulationID)
// }

// func (s *regulationService) GetIDByPseudo(ctx context.Context, pseudoId string) (uint64, error) {
// 	return s.storage.GetIDByPseudo(ctx, pseudoId)
// }
