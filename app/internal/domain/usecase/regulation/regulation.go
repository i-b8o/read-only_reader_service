package usecase_regulation

import (
	"context"
	"regulations_service/internal/domain/entity"
)

type RegulationService interface {
	CreateOne(ctx context.Context, regulation entity.Regulation) (uint64, error)
}

// type ChapterService interface {
// 	GetAllById(ctx context.Context, regulationID uint64) ([]entity.Chapter, error)
// 	GetOrderNum(ctx context.Context, id uint64) (orderNum uint64, err error)
// 	DeleteForRegulation(ctx context.Context, regulationID uint64) error
// 	GetIDByPseudo(ctx context.Context, pseudoId string) (uint64, error)
// }
// type ParagraphService interface {
// 	GetAllById(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error)
// 	UpdateOne(ctx context.Context, content string, paragraphID uint64) error
// 	DeleteForChapter(ctx context.Context, chapterID uint64) error
// }

// type SpeechService interface {
// 	GetAllById(ctx context.Context, paragraphID uint64) ([]entity.Speech, error)
// 	DeleteForParagraph(ctx context.Context, paragraphID uint64) error
// }

type regulationUsecase struct {
	regulationService RegulationService
	// chapterService    ChapterService
	// paragraphService  ParagraphService
	// speechService     SpeechService
}

func NewRegulationUsecase(regulationService RegulationService) *regulationUsecase {
	return &regulationUsecase{regulationService: regulationService}
}

func (u regulationUsecase) CreateRegulation(ctx context.Context, regulation entity.Regulation) (uint64, error) {
	return u.regulationService.CreateOne(ctx, regulation)
}

// func (u regulationUsecase) GetFullRegulationByID(ctx context.Context, regulationID uint64) (entity.Regulation, error) {
// 	regulation, err := u.regulationService.GetOne(ctx, regulationID)
// 	if err != nil {
// 		return entity.Regulation{}, err
// 	}
// 	chapters, err := u.chapterService.GetAllById(ctx, regulationID)
// 	if err != nil {
// 		return entity.Regulation{}, err
// 	}

// 	for _, chapter := range chapters {
// 		paragraphs, err := u.paragraphService.GetAllById(ctx, chapter.ID)
// 		if err != nil {
// 			return entity.Regulation{}, err
// 		}
// 		chapter.Paragraphs = paragraphs
// 	}

// 	regulation.Chapters = chapters

// 	return regulation, nil
// }

// func (u regulationUsecase) GetDartFullRegulationByID(ctx context.Context, regulationID uint64) string {
// 	regulation, err := u.regulationService.GetOne(ctx, regulationID)
// 	if err != nil {
// 		return ""
// 	}
// 	chapters, err := u.chapterService.GetAllById(ctx, regulationID)
// 	if err != nil {
// 		return ""
// 	}

// 	dartClass := `
// 	import 'paragraph.dart';
// 	import 'chapter.dart';

// 	class Regulation {
// 		static const int id = %d;
// 		static const String name = "%s";
// 		static const String abbreviation = "%s";
// 		static const List<Chapter> chapters = <Chapter>[
// 			%s
// 		];
// 	}
// 	`

// 	chaptersDartString, _ := u.chaptersDart(ctx, chapters)
// 	return fmt.Sprintf(dartClass, regulationID, regulation.Name, regulation.Abbreviation, chaptersDartString)
// }

// func (u regulationUsecase) chaptersDart(ctx context.Context, chapters []entity.Chapter) (dartChaptersString string, err error) {
// 	dartChapter := `Chapter(id: %d, name: "%s", num: "%s", orderNum: %d , paragraphs: [
// 		%s
// 	]),`
// 	for _, chapter := range chapters {
// 		paragraphs, err := u.paragraphService.GetAllById(ctx, chapter.ID)
// 		if err != nil {
// 			return dartChapter, err
// 		}
// 		dartPar := paragraphsDart(ctx, paragraphs, u)

// 		num := ""
// 		if len(chapter.Num) > 0 {
// 			num = chapter.Num
// 		}
// 		name := strings.Replace(chapter.Name, "\n", "", -1)
// 		temp := fmt.Sprintf(dartChapter, chapter.ID, name, num, chapter.OrderNum, dartPar)
// 		dartChaptersString += temp
// 	}
// 	return dartChaptersString, nil
// }

// func speachText(contentSlice []string) string {
// 	start := `[`
// 	end := `]`

// 	var result string

// 	for _, part := range contentSlice {
// 		part = strings.ReplaceAll(part, "'", `"`)
// 		str := fmt.Sprintf(`'%s',`, part)
// 		result += str
// 	}
// 	result = result[:len(result)-1]
// 	re := regexp.MustCompile(`\r?\n`)
// 	result = re.ReplaceAllString(result, " ")
// 	return start + result + end
// }

// func paragraphsDart(ctx context.Context, paragraphs []entity.Paragraph, u regulationUsecase) (dartParagraphsList string) {
// 	for _, p := range paragraphs {
// 		text := strings.Replace(p.Content, "\n", "", -1)
// 		text = strings.ReplaceAll(text, `'`, `"`)

// 		var speechTextSlice []string
// 		if p.ID > 0 {
// 			speechSlice, err := u.speechService.GetAllById(ctx, p.ID)
// 			if err != nil {
// 				return ""
// 			}
// 			for _, t := range speechSlice {
// 				speechTextSlice = append(speechTextSlice, t.Content)
// 			}
// 		} else {
// 			speechTextSlice = append(speechTextSlice, p.Content)
// 		}

// 		textToSpeech := speachText(speechTextSlice)
// 		dartParagraphsList += fmt.Sprintf(`		Paragraph(id: %d, num: %d, isTable: %t,isNFT: %t, paragraphClass: "%s", content: '%s', chapterID: %d, textToSpeech: %s),
// 		`, p.ID, p.Num, p.IsTable, p.IsNFT, p.Class, text, p.ChapterID, textToSpeech)
// 	}
// 	return dartParagraphsList
// }

// func (u regulationUsecase) GetDocumentRoot(ctx context.Context, stringID string) (entity.Regulation, []entity.Chapter) {
// 	uint64ID, err := strconv.ParseUint(stringID, 10, 64)
// 	if err != nil {
// 		return entity.Regulation{}, nil
// 	}
// 	regulation, err := u.regulationService.GetOne(ctx, uint64ID)
// 	if err != nil {
// 		return entity.Regulation{}, nil
// 	}

// 	regulation.Name = nonsense.Capitalize(regulation.Name)
// 	chapters, err := u.chapterService.GetAllById(ctx, uint64ID)
// 	if err != nil {
// 		return entity.Regulation{}, nil
// 	}
// 	return regulation, chapters
// }

// func (u regulationUsecase) GetDocuments(ctx context.Context) []entity.Regulation {
// 	regulations, err := u.regulationService.GetAll(ctx)
// 	if err != nil {
// 		return nil
// 	}

// 	return regulations
// }

// func (u regulationUsecase) DeleteRegulation(ctx context.Context, regulationID uint64) error {
// 	chapters, err := u.chapterService.GetAllById(ctx, regulationID)
// 	if err != nil {
// 		return err
// 	}

// 	for _, chapter := range chapters {

// 		paragraphs, err := u.paragraphService.GetAllById(ctx, chapter.ID)
// 		if err != nil {
// 			return err
// 		}
// 		for _, paragraph := range paragraphs {
// 			err = u.speechService.DeleteForParagraph(ctx, paragraph.ID)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		err = u.paragraphService.DeleteForChapter(ctx, chapter.ID)
// 		if err != nil {
// 			return err
// 		}

// 	}
// 	err = u.chapterService.DeleteForRegulation(ctx, regulationID)
// 	if err != nil {
// 		return err
// 	}
// 	err = u.regulationService.DeleteRegulation(ctx, regulationID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
