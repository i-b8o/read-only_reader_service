package usecase_chapter

import (
	"context"
	"regulations_service/internal/domain/entity"
)

type ChapterService interface {
	Create(ctx context.Context, chapter entity.Chapter) (string, error)
	// GetOneById(ctx context.Context, chapterID uint64) (entity.Chapter, error)
	// GetAllById(ctx context.Context, regulationID uint64) ([]entity.Chapter, error)
}

type ParagraphService interface {
	GetAllById(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error)
}

type RegulationService interface {
	GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error)
}

type chapterUsecase struct {
	chapterService    ChapterService
	paragraphService  ParagraphService
	regulationService RegulationService
}

func NewChapterUsecase(chapterService ChapterService, paragraphService ParagraphService, regulationService RegulationService) *chapterUsecase {
	return &chapterUsecase{chapterService: chapterService, paragraphService: paragraphService, regulationService: regulationService}
}

func (u chapterUsecase) CreateChapter(ctx context.Context, chapter entity.Chapter) (string, error) {
	return u.chapterService.Create(ctx, chapter)
}

// func (u chapterUsecase) GetChapter(ctx context.Context, chapterID string) (entity.Regulation, error) {
// 	uint64ID, err := strconv.ParseUint(chapterID, 10, 64)
// 	if err != nil {
// 		return entity.Regulation{}, err
// 	}

// 	chapter, err := u.chapterService.GetOneById(ctx, uint64ID)
// 	if err != nil {
// 		return entity.Regulation{}, err
// 	}

// 	chapter.Paragraphs, err = u.paragraphService.GetAllById(ctx, uint64ID)
// 	if err != nil {
// 		return entity.Regulation{}, err
// 	}
// 	regulation, err := u.regulationService.GetOne(ctx, chapter.RegulationID)
// 	if err != nil {
// 		return entity.Regulation{}, err
// 	}
// 	chapters, err := u.chapterService.GetAllById(ctx, chapter.RegulationID)
// 	if err != nil {
// 		return entity.Regulation{}, err
// 	}
// 	regulation.Chapters = chapters
// 	return regulation, chapter
// }
