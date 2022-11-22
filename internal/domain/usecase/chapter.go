package usecase

import (
	"context"
	"read-only_reader_service/internal/domain/entity"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChapterService interface {
	Get(ctx context.Context, chapterID uint64) (*entity.Chapter, error)
	GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReaderChapter, error)
}

type ParagraphService interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
}

type chapterUsecase struct {
	chapterService   ChapterService
	paragraphService ParagraphService
}

func NewChapterUsecase(chapterService ChapterService, paragraphService ParagraphService) *chapterUsecase {
	return &chapterUsecase{chapterService: chapterService, paragraphService: paragraphService}
}

func (u *chapterUsecase) Get(ctx context.Context, chapterID uint64) (*pb.GetOneChapterResponse, error) {
	chapterInfo, err := u.chapterService.Get(ctx, chapterID)
	if err != nil {
		return nil, err
	}
	paragraphs, err := u.paragraphService.GetAll(ctx, chapterID)
	if err != nil {
		return nil, err
	}
	return &pb.GetOneChapterResponse{ID: chapterID, Name: chapterInfo.Name, Num: chapterInfo.Num, RegulationID: chapterInfo.RegulationID, OrderNum: chapterInfo.OrderNum, Paragraphs: paragraphs, UpdatedAt: timestamppb.New(chapterInfo.UpdatedAt)}, nil
}

func (u *chapterUsecase) GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReaderChapter, error) {
	return u.chapterService.GetAll(ctx, regulationID)
}