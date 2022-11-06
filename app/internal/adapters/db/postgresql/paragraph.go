package postgressql

import (
	"context"
	"fmt"
	"regulations_service/internal/domain/entity"
	client "regulations_service/pkg/client/postgresql"
)

type paragraphStorage struct {
	client client.PostgreSQLClient
}

func NewParagraphStorage(client client.PostgreSQLClient) *paragraphStorage {
	return &paragraphStorage{client: client}
}

// CreateAll
func (ps *paragraphStorage) CreateAll(ctx context.Context, paragraphs []entity.Paragraph) error {
	vals := []interface{}{}
	sql := `INSERT INTO paragraphs ("paragraph_id","order_num","is_table","is_nft","has_links","class","content","c_id") VALUES `
	i := 1
	for _, p := range paragraphs {
		sql += fmt.Sprintf("($%d, $%d, $%d , $%d, $%d, $%d, $%d, $%d),", i, i+1, i+2, i+3, i+4, i+5, i+6, i+7)
		i = i + 8
		vals = append(vals, p.ID, p.Num, p.IsTable, p.IsNFT, p.HasLinks, p.Class, p.Content, p.ChapterID)
	}
	sql = sql[:len(sql)-1]

	if _, err := ps.client.Exec(ctx, sql, vals...); err != nil {
		return err
	}

	return nil
}

func (ps *paragraphStorage) UpdateOne(ctx context.Context, content string, paragraphID uint64) error {
	sql := `UPDATE "paragraphs" SET content = $1 WHERE paragraph_id = $2 RETURNING "id"`
	row := ps.client.QueryRow(ctx, sql, content, paragraphID)
	var ID uint64

	err := row.Scan(&ID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllById returns all paragraphs associated with the given chapter ID
func (ps *paragraphStorage) GetAllForChapter(ctx context.Context, chapterID uint64) ([]entity.Paragraph, error) {
	const sql = `SELECT paragraph_id, order_num, is_nft, is_table, has_links, class, content, c_id FROM "paragraphs" WHERE c_id = $1 AND content!='-' ORDER BY order_num`

	var paragraphs []entity.Paragraph

	rows, err := ps.client.Query(ctx, sql, chapterID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		paragraph := entity.Paragraph{}
		if err = rows.Scan(
			&paragraph.ID, &paragraph.Num, &paragraph.IsNFT, &paragraph.IsTable, &paragraph.HasLinks, &paragraph.Class, &paragraph.Content, &paragraph.ChapterID,
		); err != nil {
			return nil, err
		}

		paragraphs = append(paragraphs, paragraph)
	}

	return paragraphs, nil
}

// Delete
func (ps *paragraphStorage) DeleteAllForChapter(ctx context.Context, chapterID uint64) error {
	sql := `delete from paragraphs where c_id=$1`
	_, err := ps.client.Exec(ctx, sql, chapterID)
	return err
}
