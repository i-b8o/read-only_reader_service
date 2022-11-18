package postgressql

import (
	"context"
	"log"
	"read-only_reader_service/pkg/client/postgresql"
	"testing"
	"time"
)

func startRegulationTest() *regulationStorage {

	pgConfig := postgresql.NewPgConfig(
		"reader", "postgres",
		"0.0.0.0", "5436", "reader",
	)

	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		log.Fatal(err)
	}
	return NewRegulationStorage(pgClient)
}

func TestGetR(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	regulationStorage := startRegulationTest()

	tests := []struct {
		id uint64

		Name         string
		Abbreviation string
		Title        string
	}{
		{
			id:           1,
			Name:         "Имя первой записи",
			Abbreviation: "Аббревиатура первой записи",
			Title:        "Заголовок первой записи",
		},
		{
			id:   0,
			Name: "",
		},
	}

	for _, tt := range tests {
		id := tt.id
		resp, err := regulationStorage.Get(ctx, id)
		if err != nil {
			t.Errorf("TestGet(%v) got unexpected error", err)
		}
		if resp.Name == "" && tt.Name == "" {
			return
		}
		if resp.Abbreviation != tt.Abbreviation {
			t.Errorf("Get(%v)=%v, wanted Abbreviation %v", tt.id, resp.Abbreviation, tt.Abbreviation)
		}
		if resp.Name != tt.Name {
			t.Errorf("Get(%v)=%v, wanted name %v", tt.id, resp.Name, tt.Name)
		}
		if resp.Title != tt.Title {
			t.Errorf("Get(%v)=%v, wanted Title %v", tt.id, resp.Title, tt.Title)
		}
	}

}
