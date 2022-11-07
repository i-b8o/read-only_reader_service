package postgressql

import (
	"context"

	pb "regulations_service/internal/pb"
	client "regulations_service/pkg/client/postgresql"
)

type regulationStorage struct {
	client client.PostgreSQLClient
}

func NewRegulationStorage(client client.PostgreSQLClient) *regulationStorage {
	return &regulationStorage{client: client}
}

func (rs *regulationStorage) Get(ctx context.Context, regulationID uint64) (*pb.Regulation, error) {
	const sql = `SELECT name,abbreviation,title FROM "regulations" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, regulationID)

	regulation := &pb.Regulation{}
	err := row.Scan(&regulation.Name, &regulation.Abbreviation, &regulation.Title)
	if err != nil {
		return regulation, err
	}

	return regulation, nil
}
