package entity

import (
	"time"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Chapter struct {
	ID           uint64
	Name         string
	Num          string
	OrderNum     uint32
	RegulationID uint64
	UpdatedAt    time.Time
}

func ChapterToGetOneChapterResponse(c *Chapter) *pb.GetOneChapterResponse {
	updated := timestamppb.New(c.UpdatedAt)

	return &pb.GetOneChapterResponse{ID: c.ID, Name: c.Name, Num: c.Num, RegulationID: c.RegulationID, OrderNum: c.OrderNum, UpdatedAt: updated}
}
