package postgressql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"read-only_reader_service/internal/domain/entity"
	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
	"github.com/jackc/pgconn"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient) *chapterStorage {
	return &chapterStorage{client: client}
}

// Get returns an chapter associated with the given ID
func (cs *chapterStorage) Get(ctx context.Context, chapterID uint64) (*entity.Chapter, error) {
	const sql = `SELECT c.id,c.name,c.num,c.order_num,c.r_id,c.updated_at, p.paragraph_id, p.order_num,p. is_nft, p.is_table, p.has_links, p.class, p.content FROM "chapter" AS c JOIN "paragraph" AS p ON p.c_id = c.id WHERE c.id = $1 ORDER BY p.order_num`
	chapter := &entity.Chapter{}

	rows, err := cs.client.Query(ctx, sql, chapterID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
		}
		return nil, fmt.Errorf("query: %s", err)
	}
	defer rows.Close()

	var id, rId uint64
	var name, num string
	var orderNum uint32
	var updated time.Time

	var paragraphs []*entity.Paragraph
	for rows.Next() {
		p := entity.Paragraph{}

		if err = rows.Scan(
			&id, &name, &num, &orderNum, &rId, &updated, &p.ID, &p.Num, &p.IsNFT, &p.IsTable, &p.HasLinks, &p.Class, &p.Content,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
			}
			return nil, fmt.Errorf("next: %s", err)
		}

		paragraphs = append(paragraphs, &p)
	}
	chapter.Paragraphs = paragraphs
	return chapter, nil
}

// GetAll returns all chapters associated with the given ID
func (cs *chapterStorage) GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReaderChapter, error) {
	const sql = `SELECT id,name,num,order_num FROM "chapter" WHERE r_id = $1 ORDER BY order_num`

	var chapters []*pb.ReaderChapter

	rows, err := cs.client.Query(ctx, sql, regulationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chapter := &pb.ReaderChapter{}
		if err = rows.Scan(
			&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				return nil, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
			}
			return nil, err
		}

		chapters = append(chapters, chapter)
	}

	return chapters, nil
}
