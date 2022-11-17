package postgressql

import (
	"context"
	"log"
	"read-only_reader_service/pkg/client/postgresql"
	"testing"
	"time"
)

func Init() *chapterStorage {

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

func TestGet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chapterAdapter := Init()

	tests := []struct {
		id uint64

		ID       uint64
		RId      uint64
		Name     string
		OrderNum uint32
		Num      string
	}{
		{
			id:       1,
			ID:       1,
			RId:      1,
			Name:     "Имя первой записи в главы",
			OrderNum: 1,
			Num:      "I",
		},
	}

	rows, err := chapterAdapter.client.Query(ctx, "select * from paragraph")
	if err != nil {
		t.Errorf("AAA got unexpected error %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		t.Log("SSSSSSSSSSSSSSSSS")
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
		if resp.Name != tt.Name {
			t.Errorf("Get(%v)=%v, wanted name %v", tt.id, resp.Name, tt.Name)
		}
	}

}
