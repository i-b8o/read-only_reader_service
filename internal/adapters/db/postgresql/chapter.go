package postgressql

import (
	"context"

	"read-only_reader_service/internal/domain/entity"
	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"

	"github.com/jackc/pgx/v4"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient) *chapterStorage {
	return &chapterStorage{client: client}
}

// Get returns an chapter associated with the given ID
func (cs *chapterStorage) Get(ctx context.Context, chapterID uint64) (*entity.Chapter, error) {
	const sql = `SELECT id,name,num,order_num,r_id,updated_at FROM "chapter" WHERE id = $1 ORDER BY order_num`
	row := cs.client.QueryRow(ctx, sql, chapterID)
	chapter := &entity.Chapter{}
	err := row.Scan(&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum, &chapter.RegulationID, &chapter.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &entity.Chapter{ID: 0}, nil
		}
		return chapter, err
	}
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
			return nil, err
		}

		chapters = append(chapters, chapter)
	}
	return chapters, nil
}
