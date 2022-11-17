package postgressql

import (
	"context"
	"errors"
	"fmt"

	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"github.com/jackc/pgconn"
)

type regulationStorage struct {
	client client.PostgreSQLClient
}

func NewRegulationStorage(client client.PostgreSQLClient) *regulationStorage {
	return &regulationStorage{client: client}
}

func (rs *regulationStorage) Get(ctx context.Context, regulationID uint64) (*pb.GetOneRegulationResponse, error) {
	const sql = `SELECT name,abbreviation,title FROM "regulation" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, regulationID)

	regulation := &pb.GetOneRegulationResponse{}
	err := row.Scan(&regulation.Name, &regulation.Abbreviation, &regulation.Title)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return regulation, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
		}
		return regulation, err
	}

	return regulation, nil
}
