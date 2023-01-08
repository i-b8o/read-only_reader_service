package postgressql

import (
	"context"
	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type typeStorage struct {
	client client.PostgreSQLClient
}

func NewTypeStorage(client client.PostgreSQLClient) *typeStorage {
	return &typeStorage{client: client}
}

// GetAll returns all types
func (s *typeStorage) GetAll(ctx context.Context) ([]*pb.TypeResponse, error) {
	const sql = `select id, name from type order by type`

	var docTypes []*pb.TypeResponse

	rows, err := s.client.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		docType := &pb.TypeResponse{}
		if err = rows.Scan(
			&docType.ID, &docType.Name,
		); err != nil {
			return nil, err
		}

		docTypes = append(docTypes, docType)
	}
	return docTypes, nil
}
