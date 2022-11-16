package service

import (
	"context"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type RegulationStorage interface {
	Get(ctx context.Context, regulationID uint64) (*pb.GetRegulationResponse, error)
}

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (*pb.GetChapterResponse, error)
	GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReaderChapter, error)
}

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
}

//	type SearchStorage interface {
//		SearchPargaraphs(ctx context.Context, searchQuery string, params ...string) ([]*pb., error)
//		SearchChapters(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
//		SearchRegulations(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
//		Search(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
//		SearchLike(ctx context.Context, searchQuery string, params ...string) ([]*pb.SearchResponse, error)
//	}
type ReadOnlyRegulationGRPCService struct {
	regulationStorage RegulationStorage
	chapterStorage    ChapterStorage
	paragraphStorage  ParagraphStorage
	// searchStorage     SearchStorage
	logging logging.Logger
	pb.UnimplementedReaderGRPCServer
}

func NewReaderGRPCService(regulationStorage RegulationStorage, chapterStorage ChapterStorage, paragraphStorage ParagraphStorage, loging logging.Logger) *ReadOnlyRegulationGRPCService {
	return &ReadOnlyRegulationGRPCService{
		regulationStorage: regulationStorage,
		chapterStorage:    chapterStorage,
		paragraphStorage:  paragraphStorage,
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
