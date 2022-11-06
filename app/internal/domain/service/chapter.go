package service

import (
	"context"
	"regulations_service/internal/domain/entity"
)

type ChapterStorage interface {
	CreateOne(ctx context.Context, chapter entity.Chapter) (uint64, error)
	GetOneById(ctx context.Context, chapterID uint64) (entity.Chapter, error)
	// GetAllById(ctx context.Context, regulationID uint64) ([]entity.Chapter, error)
	// GetOrderNum(ctx context.Context, id uint64) (orderNum uint64, err error)
	// DeleteForRegulation(ctx context.Context, regulationID uint64) error
	// GetIDByPseudo(ctx context.Context, pseudoId string) (uint64, error)
}

type chapterService struct {
	storage ChapterStorage
}

func NewChapterService(storage ChapterStorage) *chapterService {
	return &chapterService{storage: storage}
}

func (s chapterService) CreateOne(ctx context.Context, chapter entity.Chapter) (uint64, error) {
	return s.storage.CreateOne(ctx, chapter)
}

func (s chapterService) GetOneById(ctx context.Context, chapterID uint64) (entity.Chapter, error) {
	return s.storage.GetOneById(ctx, chapterID)
}

// func (s chapterService) GetAllById(ctx context.Context, regulationID uint64) ([]entity.Chapter, error) {
// 	return s.storage.GetAllById(ctx, regulationID)
// }

// func (s chapterService) GetOrderNum(ctx context.Context, id uint64) (orderNum uint64, err error) {
// 	return s.storage.GetOrderNum(ctx, id)
// }

// func (s chapterService) DeleteForRegulation(ctx context.Context, regulationID uint64) error {
// 	return s.storage.DeleteForRegulation(ctx, regulationID)
// }

// func (s *chapterService) GetIDByPseudo(ctx context.Context, pseudoId string) (uint64, error) {
// 	return s.storage.GetIDByPseudo(ctx, pseudoId)
// }
