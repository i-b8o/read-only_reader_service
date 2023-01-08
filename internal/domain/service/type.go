package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type TypeStorage interface {
	GetAll(ctx context.Context) ([]*pb.TypeResponse, error)
}
type typeService struct {
	storage TypeStorage
	logger  logging.Logger
}

func NewTypeService(storage TypeStorage, logger logging.Logger) *typeService {
	return &typeService{storage: storage, logger: logger}
}

func (s *typeService) GetAll(ctx context.Context) (*pb.GetAllTypesResponse, error) {
	types, err := s.storage.GetAll(ctx)
	if err != nil {
		s.logger.Errorf("could not get types: %v", err)
		return nil, err
	}
	return &pb.GetAllTypesResponse{Types: types}, nil
}
