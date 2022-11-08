package service

import (
	"context"

	"github.com/i-b8o/logging"
	"github.com/i-b8o/regulations_contracts/pb"
)

type RegulationStorage interface {
	Get(ctx context.Context, regulationID uint64) (*pb.GetRegulationResponse, error)
}

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (*pb.GetChapterResponse, error)
	GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReadOnlyChapter, error)
}

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReadOnlyParagraph, error)
}

type SearchStorage interface {
	SearchPargaraphs(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
	SearchChapters(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
	SearchRegulations(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
	Search(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
	SearchLike(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
}
type ReadOnlyRegulationGRPCService struct {
	regulationStorage RegulationStorage
	chapterStorage    ChapterStorage
	paragraphStorage  ParagraphStorage
	searchStorage     SearchStorage
	logging           logging.Logger
	pb.UnimplementedReadOnlyRegulationGRPCServer
}

func NewReadOnlyRegulationGRPCService(regulationStorage RegulationStorage, chapterStorage ChapterStorage, paragraphStorage ParagraphStorage, searchStorage SearchStorage, loging logging.Logger) *ReadOnlyRegulationGRPCService {
	return &ReadOnlyRegulationGRPCService{
		regulationStorage: regulationStorage,
		chapterStorage:    chapterStorage,
		paragraphStorage:  paragraphStorage,
		searchStorage:     searchStorage,
		logging:           loging,
	}
}

func (s *ReadOnlyRegulationGRPCService) GetRegulation(ctx context.Context, req *pb.GetRegulationRequest) (*pb.GetRegulationResponse, error) {
	id := req.GetID()
	return s.regulationStorage.Get(ctx, id)
}

func (s *ReadOnlyRegulationGRPCService) GetChapter(ctx context.Context, req *pb.GetChapterRequest) (*pb.GetChapterResponse, error) {
	id := req.GetID()
	return s.chapterStorage.Get(ctx, id)
}

func (s *ReadOnlyRegulationGRPCService) GetAllChapters(ctx context.Context, req *pb.GetAllChaptersRequest) (*pb.GetAllChaptersResponse, error) {
	id := req.GetID()
	chapters, err := s.chapterStorage.GetAll(ctx, id)
	return &pb.GetAllChaptersResponse{Chapters: chapters}, err
}

func (s *ReadOnlyRegulationGRPCService) GetParagraphs(ctx context.Context, req *pb.GetParagraphsRequest) (*pb.GetParagraphsResponse, error) {
	id := req.GetID()
	paragraps, err := s.paragraphStorage.GetAll(ctx, id)
	return &pb.GetParagraphsResponse{Paragraphs: paragraps}, err
}

func (s *ReadOnlyRegulationGRPCService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponseMessage, error) {
	offset := req.GetOffset()
	limit := req.GetLimit()
	query := req.GetSearchQuery()

	res, err := s.searchStorage.Search(ctx, query, offset, limit)
	return &pb.SearchResponseMessage{Response: res}, err
}

func (s *ReadOnlyRegulationGRPCService) SearchRegulations(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponseMessage, error) {
	offset := req.GetOffset()
	limit := req.GetLimit()
	query := req.GetSearchQuery()

	res, err := s.searchStorage.SearchRegulations(ctx, query, offset, limit)
	return &pb.SearchResponseMessage{Response: res}, err
}

func (s *ReadOnlyRegulationGRPCService) SearchChapters(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponseMessage, error) {
	offset := req.GetOffset()
	limit := req.GetLimit()
	query := req.GetSearchQuery()

	res, err := s.searchStorage.SearchChapters(ctx, query, offset, limit)
	return &pb.SearchResponseMessage{Response: res}, err
}

func (s *ReadOnlyRegulationGRPCService) SearchPargaraphs(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponseMessage, error) {
	offset := req.GetOffset()
	limit := req.GetLimit()
	query := req.GetSearchQuery()

	res, err := s.searchStorage.SearchPargaraphs(ctx, query, offset, limit)
	return &pb.SearchResponseMessage{Response: res}, err
}
