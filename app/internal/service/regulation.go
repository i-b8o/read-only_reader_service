package service

import (
	"context"

	"regulations_read_only_service/internal/pb"

	"github.com/i-b8o/logging"
)

type RegulationStorage interface {
	Get(ctx context.Context, regulationID uint64) (*pb.Regulation, error)
}

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (*pb.Chapter, error)
	GetAll(ctx context.Context, regulationID uint64) ([]*pb.Chapter, error)
}

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.Paragraph, error)
}

type SearchStorage interface {
	Search(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResult, error)
	SearchLike(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResult, error)
	SearchRegulations(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResult, error)
	SearchChapters(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResult, error)
	SearchPargaraphs(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResult, error)
}
type RegulationGRPCService struct {
	regulationStorage RegulationStorage
	chapterStorage    ChapterStorage
	paragraphStorage  ParagraphStorage
	searchStorage     SearchStorage
	logging           logging.Logger
	pb.UnimplementedRegulationGRPCServer
}

func NewRegulationGRPCService(regulationStorage RegulationStorage, chapterStorage ChapterStorage, paragraphStorage ParagraphStorage, searchStorage SearchStorage, loging logging.Logger) *RegulationGRPCService {
	return &RegulationGRPCService{
		regulationStorage: regulationStorage,
		chapterStorage:    chapterStorage,
		paragraphStorage:  paragraphStorage,
		searchStorage:     searchStorage,
		logging:           loging,
	}
}

func (s *RegulationGRPCService) GetRegulation(ctx context.Context, req *pb.ID) (*pb.Regulation, error) {
	id := req.GetID()
	return s.regulationStorage.Get(ctx, id)
}

func (s *RegulationGRPCService) GetChapter(ctx context.Context, req *pb.ID) (*pb.Chapter, error) {
	return &pb.Chapter{}, nil
}

func (s *RegulationGRPCService) GetAllChapters(ctx context.Context, req *pb.ID) (*pb.Chapters, error) {
	return &pb.Chapters{}, nil
}

func (s *RegulationGRPCService) GetParagraphs(ctx context.Context, req *pb.ID) (*pb.Paragraphs, error) {
	return &pb.Paragraphs{}, nil
}

func (s *RegulationGRPCService) Search(ctx context.Context, req *pb.SearchRequestMessage) (*pb.SearchResponseMessage, error) {
	return &pb.SearchResponseMessage{}, nil
}

func (s *RegulationGRPCService) SearchRegulations(ctx context.Context, req *pb.SearchRequestMessage) (*pb.SearchResponseMessage, error) {
	return &pb.SearchResponseMessage{}, nil
}

func (s *RegulationGRPCService) SearchChapters(ctx context.Context, req *pb.SearchRequestMessage) (*pb.SearchResponseMessage, error) {
	return &pb.SearchResponseMessage{}, nil
}

func (s *RegulationGRPCService) SearchPargaraphs(ctx context.Context, req *pb.SearchRequestMessage) (*pb.SearchResponseMessage, error) {
	return &pb.SearchResponseMessage{}, nil
}
