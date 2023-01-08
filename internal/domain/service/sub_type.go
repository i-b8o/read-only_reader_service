package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type SubTypeStorage interface {
	GetAll(ctx context.Context, typeID uint64) ([]*pb.SubtypeResponse, error)
}
type subTypeService struct {
	storage SubTypeStorage
	logger  logging.Logger
}

func NewSubTypeService(storage SubTypeStorage, logger logging.Logger) *subTypeService {
	return &subTypeService{storage: storage, logger: logger}
}

func (s *subTypeService) GetAll(ctx context.Context, typeID uint64) (*pb.GetAllSubtypesResponse, error) {
	subTypes, err := s.storage.GetAll(ctx, typeID)
	if err != nil {
		s.logger.Errorf("could not get subtypes for %d, %v", typeID, err)
		return nil, err
	}
	return &pb.GetAllSubtypesResponse{Subtypes: subTypes}, nil
}
