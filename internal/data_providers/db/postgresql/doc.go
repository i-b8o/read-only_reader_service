package postgressql

import (
	"context"
	"errors"
	"fmt"

	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"github.com/jackc/pgconn"
)

type docStorage struct {
	client client.PostgreSQLClient
}

func NewDocStorage(client client.PostgreSQLClient) *docStorage {
	return &docStorage{client: client}
}

func (rs *docStorage) Get(ctx context.Context, docID uint64) (*pb.GetOneDocResponse, error) {
	const sql = `SELECT doc.name, chapter.id, chapter.num,chapter.order_num,chapter.name FROM "doc" LEFT JOIN chapter ON doc.id = chapter.doc_id where doc.id=$1`

	var chapters []*pb.Chapter
	var docName string

	rows, err := rs.client.Query(ctx, sql, docID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		chapter := &pb.Chapter{}
		if err = rows.Scan(
			&docName, &chapter.ID, &chapter.Num, &chapter.OrderNum, &chapter.Name,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
			}
			return nil, err
		}
		fmt.Println(chapter.Name)
		chapters = append(chapters, chapter)
	}

	return &pb.GetOneDocResponse{Name: docName, Chapters: chapters}, nil

}
