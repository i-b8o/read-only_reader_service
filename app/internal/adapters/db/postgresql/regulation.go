package postgressql

import (
	"context"
	"regulations_service/internal/domain/entity"
	"time"

	client "regulations_service/pkg/client/postgresql"
)

type regulationStorage struct {
	client client.PostgreSQLClient
}

func NewRegulationStorage(client client.PostgreSQLClient) *regulationStorage {
	return &regulationStorage{client: client}
}

// GetOne returns regulation associated with the given ID
// Create returns the ID of the inserted chapter
func (rs *regulationStorage) CreateOne(ctx context.Context, regulation entity.Regulation) (uint64, error) {
	t := time.Now()

	const sql = `INSERT INTO regulations ("name", "abbreviation", "title", "created_at") VALUES ($1, $2, $3, $4) RETURNING "id"`

	row := rs.client.QueryRow(ctx, sql, regulation.Name, regulation.Abbreviation, regulation.Title, t)
	var regulationID uint64

	err := row.Scan(&regulationID)
	return regulationID, err
}

func (rs *regulationStorage) GetOne(ctx context.Context, regulationID uint64) (entity.Regulation, error) {
	const sql = `SELECT name,abbreviation,title FROM "regulations" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, regulationID)

	var regulation entity.Regulation
	err := row.Scan(&regulation.Name, &regulation.Abbreviation, &regulation.Title)
	if err != nil {
		return regulation, err
	}

	return regulation, nil
}

// Delete
func (rs *regulationStorage) DeleteRegulation(ctx context.Context, regulationID uint64) error {
	sql := `delete from regulations where id=$1`
	_, err := rs.client.Exec(ctx, sql, regulationID)

	return err
}
