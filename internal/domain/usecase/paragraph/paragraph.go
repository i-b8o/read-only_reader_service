package usecase_paragraph

import (
	"context"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type ParagraphService interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
}

type paragraphUsecase struct {
	paragraphService ParagraphService
}

func NewParagraphUsecase(paragraphService ParagraphService) *paragraphUsecase {
	return &paragraphUsecase{paragraphService: paragraphService}
}
func (u *paragraphUsecase) GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error) {
	return u.paragraphService.GetAll(ctx, chapterID)
}
