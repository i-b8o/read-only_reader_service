package service

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type DocStorage interface {
	Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error)
}

type docService struct {
	storage DocStorage
}

func NewDocService(storage DocStorage) *docService {
	return &docService{storage: storage}
}

func (s *docService) Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error) {
	return s.storage.Get(ctx, docID)
}
