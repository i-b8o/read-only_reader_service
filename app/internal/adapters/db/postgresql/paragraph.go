package postgressql

import (
	"context"
	"errors"
	"fmt"

	client "regulations_read_only_service/pkg/client/postgresql"

	"github.com/i-b8o/regulations_contracts/pb"
	"github.com/jackc/pgconn"
)

type paragraphStorage struct {
	client client.PostgreSQLClient
}

func NewParagraphStorage(client client.PostgreSQLClient) *paragraphStorage {
	return &paragraphStorage{client: client}
}

// GetAllById returns all paragraphs associated with the given chapter ID
func (ps *paragraphStorage) GetAll(ctx context.Context, chapterID uint64) ([]*pb.ReadOnlyParagraph, error) {
	const sql = `SELECT paragraph_id, order_num, is_nft, is_table, has_links, class, content, c_id FROM "paragraphs" WHERE c_id = $1 AND content!='-' ORDER BY order_num`

	var paragraphs []*pb.ReadOnlyParagraph

	rows, err := ps.client.Query(ctx, sql, chapterID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		paragraph := &pb.ReadOnlyParagraph{}
		if err = rows.Scan(
			&paragraph.ID, &paragraph.Num, &paragraph.IsNFT, &paragraph.IsTable, &paragraph.HasLinks, &paragraph.Class, &paragraph.Content, &paragraph.ChapterID,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
			}
			return nil, err
		}

		paragraphs = append(paragraphs, paragraph)
	}

	return paragraphs, nil
}
