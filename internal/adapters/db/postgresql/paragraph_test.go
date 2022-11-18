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

func startParagraphTest() *paragraphStorage {

	pgConfig := postgresql.NewPgConfig(
		"reader", "postgres",
		"0.0.0.0", "5436", "reader",
	)

	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		log.Fatal(err)
	}
	return NewParagraphStorage(pgClient)
}

func TestGetAllP(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	paragraphAdapter := startParagraphTest()

	tests := []struct {
		id       uint64
		chapters []*pb.ReaderParagraph
	}{
		{
			1,
			[]*pb.ReaderParagraph{
				&pb.ReaderParagraph{ID: 1, Num: 1, IsNFT: false, IsTable: false, HasLinks: false, Class: "any-class", Content: "Содержимое первого параграфа", ChapterID: 1},
				&pb.ReaderParagraph{ID: 2, Num: 2, IsNFT: true, IsTable: true, HasLinks: true, Class: "any-class", Content: "Содержимое второго параграфа", ChapterID: 1},
				&pb.ReaderParagraph{ID: 3, Num: 3, IsNFT: false, IsTable: false, HasLinks: false, Class: "any-class", Content: "Содержимое третьего параграфа", ChapterID: 1},
			},
		},
		{
			0,
			nil,
		},
	}
	for _, tt := range tests {
		id := tt.id
		resp, err := paragraphAdapter.GetAll(ctx, id)
		if err != nil {
			t.Errorf("TestGet(%v) got unexpected error", err)
		}

		if !reflect.DeepEqual(resp, tt.chapters) {
			t.Errorf("Get(%v)=%v, wanted %v", tt.id, resp, tt.chapters)
		}

	}
}
