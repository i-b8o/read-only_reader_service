package postgressql

import (
	"context"
	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type subTypeStorage struct {
	client client.PostgreSQLClient
}

func NewSubTypeStorage(client client.PostgreSQLClient) *subTypeStorage {
	return &subTypeStorage{client: client}
}

func (s *subTypeStorage) GetAll(ctx context.Context, typeID uint64) ([]*pb.SubtypeResponse, error) {
	const sql = `SELECT id, name FROM subtype WHERE type_id=$1`

	var docSubTypes []*pb.SubtypeResponse

	rows, err := s.client.Query(ctx, sql, typeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		docSubType := &pb.SubtypeResponse{}
		if err = rows.Scan(
			&docSubType.ID, &docSubType.Name,
		); err != nil {
			return nil, err
		}

		docSubTypes = append(docSubTypes, docSubType)
	}
	return docSubTypes, nil
}
