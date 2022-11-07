package postgressql

import (
	"context"

	"regulations_service/internal/pb"
	client "regulations_service/pkg/client/postgresql"
)

type paragraphStorage struct {
	client client.PostgreSQLClient
}

func NewParagraphStorage(client client.PostgreSQLClient) *paragraphStorage {
	return &paragraphStorage{client: client}
}

// GetAllById returns all paragraphs associated with the given chapter ID
func (ps *paragraphStorage) GetAll(ctx context.Context, chapterID uint64) ([]*pb.Paragraph, error) {
	const sql = `SELECT paragraph_id, order_num, is_nft, is_table, has_links, class, content, c_id FROM "paragraphs" WHERE c_id = $1 AND content!='-' ORDER BY order_num`

	var paragraphs []*pb.Paragraph

	rows, err := ps.client.Query(ctx, sql, chapterID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		paragraph := &pb.Paragraph{}
		if err = rows.Scan(
			&paragraph.ID, &paragraph.Num, &paragraph.IsNFT, &paragraph.IsTable, &paragraph.HasLinks, &paragraph.Class, &paragraph.Content, &paragraph.ChapterID,
		); err != nil {
			return nil, err
		}

		paragraphs = append(paragraphs, paragraph)
	}

	return paragraphs, nil
}
