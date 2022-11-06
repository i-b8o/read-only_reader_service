package grpc_controller

import (
	"context"
	"regulations_service/internal/domain/entity"

	"regulations_service/internal/pb"

	"github.com/i-b8o/logging"
)

type RegulationUsecase interface {
	CreateRegulation(ctx context.Context, regulation entity.Regulation) (uint64, error)
}

type ChapterUsecase interface {
	CreateChapter(ctx context.Context, chapter entity.Chapter) string
}

type ParagraphUsecase interface {
	CreateParagraphs(ctx context.Context, paragraphs []entity.Paragraph) error
}

type RegulationGRPCService struct {
	regulationUsecase RegulationUsecase
	logging           logging.Logger
	chapterUsecase    ChapterUsecase
	paragraphUsecase  ParagraphUsecase
	pb.UnimplementedRegulationGRPCServer
}

func NewRegulationGRPCService(regulationUsecase RegulationUsecase, chapterUsecase ChapterUsecase, paragraphUsecase ParagraphUsecase, loging logging.Logger) *RegulationGRPCService {
	return &RegulationGRPCService{
		regulationUsecase: regulationUsecase,
		chapterUsecase:    chapterUsecase,
		paragraphUsecase:  paragraphUsecase,
		logging:           loging,
	}
}

func (s *RegulationGRPCService) CreateRegulation(ctx context.Context, req *pb.CreateRegulationRequestMessage) (*pb.CreateRegulationResponseMessage, error) {
	// MAPPING
	regulation := entity.Regulation{
		Name:         req.RegulationName,
		Abbreviation: req.Abbreviation,
		Title:        req.Title,
	}
	// Usecase
	id, err := s.regulationUsecase.CreateRegulation(ctx, regulation)
	if err != nil {
		s.logging.Error(err)
		return nil, err
	}
	return &pb.CreateRegulationResponseMessage{ID: id}, nil
}

func (s *RegulationGRPCService) CreateChapter(ctx context.Context, req *pb.CreateChapterRequestMessage) (*pb.CreateChapterResponseMessage, error) {
	// MAPPING
	chapter := entity.Chapter{
		ID:           req.ChapterID,
		Name:         req.ChapterName,
		Num:          req.ChapterNum,
		RegulationID: req.RegulationID,
		OrderNum:     req.OrderNum,
	}
	// Usecase
	id := s.chapterUsecase.CreateChapter(ctx, chapter)
	return &pb.CreateChapterResponseMessage{ID: id}, nil
}

func (s *RegulationGRPCService) CreateParagraphs(ctx context.Context, req *pb.CreateParagraphsRequestMessage) (*pb.CreateParagraphsResponseMessage, error) {
	var paragraphs []entity.Paragraph
	// MAPPING
	for _, p := range req.Paragraphs {
		paragraph := entity.Paragraph{
			ID:        p.ParagraphID,
			Num:       p.ParagraphOrderNum,
			IsTable:   p.IsTable,
			IsNFT:     p.IsNFT,
			HasLinks:  p.HasLinks,
			Class:     p.ParagraphClass,
			Content:   p.ParagraphText,
			ChapterID: p.ChapterID,
		}

		if p.ParagraphID > 0 {
			paragraph.ID = p.ParagraphID
		}
		paragraphs = append(paragraphs, paragraph)
	}
	// Usecase
	err := s.paragraphUsecase.CreateParagraphs(ctx, paragraphs)
	if err != nil {
		return &pb.CreateParagraphsResponseMessage{Status: "not"}, err
	}
	return &pb.CreateParagraphsResponseMessage{Status: "ok"}, nil
}

// func (s *RegulationGRPCService) GenerateLinks(ctx context.Context, req *pb.GenerateLinksRequest) (*emptypb.Empty, error) {
// 	err := s.regulationUsecase.GenerateLinks(ctx, req.ID)
// 	return &emptypb.Empty{}, err
// }

// func (s *RegulationGRPCService) CreateChapter(ctx context.Context, req *pb.CreateChapterRequest) (*pb.CreateChapterResponse, error) {
// 	// MAPPING
// 	chapter := entity.Chapter{
// 		ID:           req.ChapterID,
// 		Pseudo:       req.PseudoID,
// 		Name:         req.ChapterName,
// 		Num:          req.ChapterNum,
// 		RegulationID: req.RegulationID,
// 		OrderNum:     req.OrderNum,
// 	}
// 	// Usecase
// 	id := s.chapterUsecase.CreateChapter(ctx, chapter)
// 	return &pb.CreateChapterResponse{Id: id}, nil
// }
