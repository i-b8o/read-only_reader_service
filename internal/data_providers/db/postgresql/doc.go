package postgressql

import (
	"context"

	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type docStorage struct {
	client client.PostgreSQLClient
}

func NewDocStorage(client client.PostgreSQLClient) *docStorage {
	return &docStorage{client: client}
}

func (rs *docStorage) Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error) {
	const sql = `SELECT name,abbreviation,title FROM "doc" WHERE id = $1 LIMIT 1`
	row := rs.client.QueryRow(ctx, sql, docID)

	doc := &pb.GetOneDocResponse{}
	err := row.Scan(&doc.Name, &doc.Abbreviation, &doc.Title)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
