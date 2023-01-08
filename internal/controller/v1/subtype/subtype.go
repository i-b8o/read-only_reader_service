package subtype_controller

import (
	"context"
	"fmt"

	"github.com/i-b8o/logging"
	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubtypeAdapter interface {
	GetAll(ctx context.Context, typeID uint64) ([]*pb.SubtypeResponse, error)
}
type SubtypeDocAdapter interface {
	GetAll(ctx context.Context, subtypeID uint64) ([]*pb.Doc, error)
}
type SubtypeGRPCService struct {
	subtypeAdapter    SubtypeAdapter
	subtypeDocAdapter SubtypeDocAdapter
	logger            logging.Logger
	pb.UnimplementedSubGRPCServer
}

func NewSubtypeGRPCService(subtypeAdapter SubtypeAdapter, subtypeDocAdapter SubtypeDocAdapter, logger logging.Logger) *SubtypeGRPCService {
	return &SubtypeGRPCService{
		subtypeAdapter:    subtypeAdapter,
		subtypeDocAdapter: subtypeDocAdapter,
		logger:            logger,
	}
}

func (s *SubtypeGRPCService) GetAll(ctx context.Context, req *pb.GetAllSubtypesRequest) (*pb.GetAllSubtypesResponse, error) {
	ID := req.GetID()
	resp, err := s.subtypeAdapter.GetAll(ctx, ID)
	if err != nil {
		s.logger.Errorf("could not get subtypes for %d, %v", ID, err)
		err := status.Errorf(codes.NotFound, fmt.Sprintf("could not get subtypes for %d, %v", ID, err))
		return nil, err
	}
	return &pb.GetAllSubtypesResponse{Subtypes: resp}, nil
}

func (s *SubtypeGRPCService) GetDocs(ctx context.Context, req *pb.GetDocsRequest) (*pb.GetDocsResponse, error) {
	ID := req.GetID()
	resp, err := s.subtypeDocAdapter.GetAll(ctx, ID)
	if err != nil {
		s.logger.Errorf("could not get docs for %d, %v", ID, err)
		err := status.Errorf(codes.NotFound, fmt.Sprintf("could not get docs for ID: %d: %v", ID, err))
		return nil, err
	}
	return &pb.GetDocsResponse{Docs: resp}, nil
}
