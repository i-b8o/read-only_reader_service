package service

import (
	"context"
	"regulations_service/internal/domain/entity"
)

type SearchStorage interface {
	Search(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	SearchLike(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	SearchChapters(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	SearchRegulations(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
	SearchPargaraphs(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error)
}

type searchService struct {
	storage SearchStorage
}

func NewSearchService(storage SearchStorage) *searchService {
	return &searchService{storage: storage}
}

func (ss searchService) Search(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.Search(ctx, searchQuery, params...)
}

func (ss searchService) SearchLike(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.SearchLike(ctx, searchQuery, params...)
}

func (ss searchService) SearchRegulations(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.SearchRegulations(ctx, searchQuery, params...)
}

func (ss searchService) SearchChapters(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.SearchChapters(ctx, searchQuery, params...)
}

func (ss searchService) SearchPargaraphs(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	return ss.storage.SearchPargaraphs(ctx, searchQuery, params...)
}
