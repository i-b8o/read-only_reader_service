package entity

import (
	"time"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Paragraph struct {
	ID        uint64
	Num       uint32
	HasLinks  bool
	IsTable   bool
	IsNFT     bool
	Class     string
	Content   string
	ChapterID uint64
}

type Chapter struct {
	ID           uint64
	Name         string
	Num          string
	OrderNum     uint32
	RegulationID uint64
	Paragraphs   []*Paragraph
	UpdatedAt    time.Time
}

func ChapterToGetOneChapterResponse(c *Chapter) *pb.GetOneChapterResponse {
	updated := timestamppb.New(c.UpdatedAt)
	paragraps := []*pb.ReaderParagraph{}
	for _, p := range c.Paragraphs {
		paragrap := &pb.ReaderParagraph{ID: p.ID, Num: p.Num, HasLinks: p.HasLinks, IsTable: p.IsTable, IsNFT: p.IsNFT, Class: p.Class, Content: p.Content, ChapterID: p.ChapterID}
		paragraps = append(paragraps, paragrap)
	}
	return &pb.GetOneChapterResponse{ID: c.ID, Name: c.Name, Num: c.Num, RegulationID: c.RegulationID, OrderNum: c.OrderNum, UpdatedAt: updated, Paragraphs: paragraps}
}
