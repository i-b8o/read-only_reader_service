package postgressql

import (
	"context"

	client "regulations_read_only_service/pkg/client/postgresql"

	"github.com/i-b8o/regulations_contracts/pb"
)

type regulationStorage struct {
	client client.PostgreSQLClient
}

func NewRegulationStorage(client client.PostgreSQLClient) *regulationStorage {
	return &regulationStorage{client: client}
}

func (rs *regulationStorage) Get(ctx context.Context, regulationID uint64) (*pb.GetRegulationResponse, error) {
	const sql = `SELECT name,abbreviation,title FROM "regulations" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, regulationID)

	regulation := &pb.GetRegulationResponse{}
	err := row.Scan(&regulation.Name, &regulation.Abbreviation, &regulation.Title)
	if err != nil {
		return regulation, err
	}

	return regulation, nil
}
