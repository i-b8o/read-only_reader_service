package usecase_regulation

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type RegulationService interface {
	Get(ctx context.Context, regulationID uint64) (*pb.GetOneRegulationResponse, error)
}

type regulationUsecase struct {
	regulationService RegulationService
}

func NewRegulationUsecase(regulationService RegulationService) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService}
}

func (u *regulationUsecase) Get(ctx context.Context, regulationID uint64) (*pb.GetOneRegulationResponse, error) {
	return u.regulationService.Get(ctx, regulationID)
}
