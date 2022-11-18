package postgressql

import (
	"context"
	"log"
	"read-only_reader_service/pkg/client/postgresql"
	"reflect"
	"testing"
	"time"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

func startChapterTest() *chapterStorage {

	pgConfig := postgresql.NewPgConfig(
		"reader", "postgres",
		"0.0.0.0", "5436", "reader",
	)

	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		log.Fatal(err)
	}
	return NewChapterStorage(pgClient)
}

func TestGetCh(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chapterAdapter := startChapterTest()

	tests := []struct {
		id uint64

		ID        uint64
		RId       uint64
		Name      string
		OrderNum  uint32
		Num       string
		UpdatedAt string
	}{
		{
			id:        1,
			ID:        1,
			RId:       1,
			Name:      "Имя первой записи",
			OrderNum:  1,
			Num:       "I",
			UpdatedAt: "2023-01-01 03:00:00 +0300 MSK",
		},
		{
			id: 0,
			ID: 0,
		},
		{
			id: 9999999999999999999,
			ID: 0,
		},
	}

	for _, tt := range tests {
		id := tt.id
		resp, err := chapterAdapter.Get(ctx, id)
		if err != nil {
			t.Errorf("TestGet(%v) got unexpected error", err)
		}
		if resp.ID != tt.ID {
			t.Errorf("Get(%v)=%v, wanted id %v", tt.id, resp.ID, tt.ID)

		}
		if resp.ID == tt.ID && resp.ID == 0 {
			return
		}
		if resp.RegulationID != tt.RId {
			t.Errorf("Get(%v)=%v, wanted r_id %v", tt.id, resp.RegulationID, tt.RId)
		}
		if resp.Name != tt.Name {
			t.Errorf("Get(%v)=%v, wanted name %v", tt.id, resp.Name, tt.Name)
		}
		if resp.OrderNum != tt.OrderNum {
			t.Errorf("Get(%v)=%v, wanted order_num %v", tt.id, resp.OrderNum, tt.OrderNum)
		}
		if resp.Num != tt.Num {
			t.Errorf("Get(%v)=%v, wanted num %v", tt.id, resp.Num, tt.Num)
		}
		if resp.UpdatedAt.String() != tt.UpdatedAt {
			t.Errorf("Get(%v)=%v, wanted updated_at %v", tt.id, resp.UpdatedAt, tt.UpdatedAt)
		}
	}

}

func TestGetAllCh(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chapterAdapter := startChapterTest()

	tests := []struct {
		id       uint64
		chapters []*pb.ReaderChapter
	}{
		{
			1,
			[]*pb.ReaderChapter{
				&pb.ReaderChapter{ID: 1, Name: "Имя первой записи", Num: "I", OrderNum: 1},
				&pb.ReaderChapter{ID: 2, Name: "Имя второй записи", Num: "II", OrderNum: 2},
				&pb.ReaderChapter{ID: 3, Name: "Имя третьей записи", Num: "III", OrderNum: 3},
			},
		},
		{
			0,
			nil,
		},
	}
	for _, tt := range tests {
		id := tt.id
		resp, err := chapterAdapter.GetAll(ctx, id)
		if err != nil {
			t.Errorf("TestGet(%v) got unexpected error", err)
		}

		if !reflect.DeepEqual(resp, tt.chapters) {
			t.Errorf("Get(%v)=%v, wanted %v", tt.id, resp, tt.chapters)
		}

	}
}
