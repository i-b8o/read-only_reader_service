package postgressql

import (
	"context"
	"read-only_reader_service/pkg/client/postgresql"
	"testing"
	"time"
)

// import (
// 	"testing"

// 	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
// )

func TestGet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pgConfig := postgresql.NewPgConfig(
		"reader", "postgres",
		"0.0.0.0", "5436", "reader",
	)

	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		t.Fatal(err)
	}
	chapterAdapter := NewChapterStorage(pgClient)

	tests := []struct {
		id                      uint64
		ID                      uint64
		RId                     uint64
		Name                    string
		OrderNum                uint32
		Num                     string
		FirstParagraphID        uint64
		SecondParagraphID       uint64
		ThirdParagraphID        uint64
		FirstParagraphNum       uint32
		SecondParagraphNum      uint32
		ThirdParagraphNum       uint32
		FirstParagraphHasLinks  bool
		SecondParagraphHasLinks bool
		ThirdParagraphHasLinks  bool
		FirstParagraphIsTable   bool
		SecondParagraphIsTable  bool
		ThirdParagraphIsTable   bool
		FirstParagraphIsNFT     bool
		SecondParagraphIsNFT    bool
		ThirdParagraphIsNFT     bool
		FirstParagraphClass     string
		SecondParagraphClass    string
		ThirdParagraphClass     string
		FirstParagraphContent   string
		SecondParagraphContent  string
		ThirdParagraphContent   string
	}{
		{
			id:                      1,
			ID:                      1,
			RId:                     1,
			Name:                    "Имя первой записи в главы",
			OrderNum:                1,
			Num:                     "I",
			FirstParagraphID:        1,
			SecondParagraphID:       2,
			ThirdParagraphID:        3,
			FirstParagraphNum:       1,
			SecondParagraphNum:      2,
			ThirdParagraphNum:       3,
			FirstParagraphHasLinks:  false,
			SecondParagraphHasLinks: true,
			ThirdParagraphHasLinks:  false,
			FirstParagraphIsTable:   false,
			SecondParagraphIsTable:  true,
			ThirdParagraphIsTable:   false,
			FirstParagraphIsNFT:     false,
			SecondParagraphIsNFT:    true,
			ThirdParagraphIsNFT:     false,
			FirstParagraphClass:     "any-class",
			SecondParagraphClass:    "any-class",
			ThirdParagraphClass:     "any-class",
			FirstParagraphContent:   "Содержимое первого параграфа",
			SecondParagraphContent:  "Содержимое второго параграфа",
			ThirdParagraphContent:   "Содержимое третьего параграфа",
		},
		{
			id: 2,
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
			t.Errorf("GetRegulation(%v)=%v, wanted id %v", tt.id, resp.ID, tt.ID)
		}
		if resp.Name != tt.Name {
			t.Errorf("GetRegulation(%v)=%v, wanted name %v", tt.id, resp.Name, tt.Name)
		}
	}

}
