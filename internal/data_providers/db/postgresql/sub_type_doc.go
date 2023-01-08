package postgressql

import (
	"context"
	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type subTypeDocStorage struct {
	client client.PostgreSQLClient
}

func NewSubTypeDocStorage(client client.PostgreSQLClient) *subTypeDocStorage {
	return &subTypeDocStorage{client: client}
}

func (s *subTypeDocStorage) GetAll(ctx context.Context, subtypeID uint64) ([]*pb.Doc, error) {
	const sql = `SELECT s.doc_id, name  FROM subtype_doc AS s INNER JOIN doc ON doc.id = s.doc_id WHERE subtype_id=$1`

	var docs []*pb.Doc

	rows, err := s.client.Query(ctx, sql, subtypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var doc pb.Doc
		if err = rows.Scan(
			&doc.ID, &doc.Name,
		); err != nil {
			return nil, err
		}

		docs = append(docs, &doc)
	}
	return docs, nil
}
