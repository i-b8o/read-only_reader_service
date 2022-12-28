package usecase_doc

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

// TODO drop usecase layer
type DocService interface {
	Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error)
}

type docUsecase struct {
	docService DocService
}

func NewDocUsecase(docService DocService) *docUsecase {
	return &docUsecase{docService: docService}
}

func (u *docUsecase) Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error) {
	return u.docService.Get(ctx, docID)
}
