package entity

import (
	"testing"
	"time"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestChapterToGetOneChapterResponse(t *testing.T) {
	d := time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC)
	dd := timestamppb.New(d)
	tests := []struct {
		chapter Chapter
		resp    *pb.GetOneChapterResponse
	}{
		{
			Chapter{ID: 1, Name: "Name", Num: "Num", OrderNum: 1, RegulationID: 1, UpdatedAt: d},
			&pb.GetOneChapterResponse{ID: 1, Name: "Name", Num: "Num", OrderNum: 1, RegulationID: 1, UpdatedAt: dd},
		},
	}
	for _, tt := range tests {
		resp := ChapterToGetOneChapterResponse(&tt.chapter)
		if resp.ID != tt.resp.ID {
			t.Errorf("ChapterToGetOneChapterResponse()=%v, wanted id %v", resp.ID, tt.resp.ID)
		}
		if resp.RegulationID != tt.resp.RegulationID {
			t.Errorf("ChapterToGetOneChapterResponse()=%v, wanted r_id %v", resp.RegulationID, tt.resp.RegulationID)
		}
		if resp.Name != tt.resp.Name {
			t.Errorf("ChapterToGetOneChapterResponse()=%v, wanted name %v", resp.Name, tt.resp.Name)
		}
		if resp.Num != tt.resp.Num {
			t.Errorf("ChapterToGetOneChapterResponse()=%v, wanted num %v", resp.Num, tt.resp.Num)
		}
		if resp.OrderNum != tt.resp.OrderNum {
			t.Errorf("ChapterToGetOneChapterResponse()=%v, wanted OrderNum %v", resp.OrderNum, tt.resp.OrderNum)
		}
		if resp.UpdatedAt != tt.resp.UpdatedAt {
			t.Errorf("ChapterToGetOneChapterResponse()=%v, wanted UpdatedAt %v", resp.UpdatedAt, tt.resp.UpdatedAt)
		}
	}
}
