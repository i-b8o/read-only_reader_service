package service

import (
	"context"
	"read-only_reader_service/internal/domain/entity"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegulationStorage interface {
	Get(ctx context.Context, regulationID uint64) (*pb.GetOneRegulationResponse, error)
}

type ChapterStorage interface {
	Get(ctx context.Context, chapterID uint64) (*entity.Chapter, error)
	GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReaderChapter, error)
}

type ParagraphStorage interface {
	GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReaderParagraph, error)
}

type ReadOnlyRegulationGRPCService struct {
	regulationStorage RegulationStorage
	chapterStorage    ChapterStorage
	logging           logging.Logger
	paragraphStorage  ParagraphStorage
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

func (s *ReadOnlyRegulationGRPCService) GetRegulation(ctx context.Context, req *pb.GetOneRegulationRequest) (*pb.GetOneRegulationResponse, error) {
	id := req.GetID()
	resp, err := s.regulationStorage.Get(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, "id was not found: %s"+err.Error())
		return nil, err
	}
	return resp, nil
}

func (s *ReadOnlyRegulationGRPCService) GetChapter(ctx context.Context, req *pb.GetOneChapterRequest) (*pb.GetOneChapterResponse, error) {
	id := req.GetID()
	chapter, err := s.chapterStorage.Get(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, "id was not found: %s"+err.Error())
		return nil, err
	}

	return entity.ChapterToGetOneChapterResponse(chapter), nil
}

func (s *ReadOnlyRegulationGRPCService) GetAllChapters(ctx context.Context, req *pb.GetAllChaptersByRegulationIdRequest) (*pb.GetAllChaptersByRegulationIdResponse, error) {
	id := req.GetID()
	chapters, err := s.chapterStorage.GetAll(ctx, id)
	if err != nil {
		err := status.Errorf(codes.NotFound, "id was not found: %s"+err.Error())
		return nil, err
	}
	return &pb.GetAllChaptersByRegulationIdResponse{Chapters: chapters}, nil
}
