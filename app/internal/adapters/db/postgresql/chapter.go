package postgressql

import (
	"context"
	"errors"
	"fmt"

	client "regulations_read_only_service/pkg/client/postgresql"

	pb "github.com/i-b8o/regulations_contracts/pb/read_only/v1"
	"github.com/jackc/pgconn"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient) *chapterStorage {
	return &chapterStorage{client: client}
}

// Get returns an chapter associated with the given ID
func (cs *chapterStorage) Get(ctx context.Context, chapterID uint64) (*pb.GetChapterResponse, error) {
	const sql = `SELECT id,name,num,order_num,r_id,updated_at FROM "chapters" WHERE id = $1 ORDER BY order_num`
	row := cs.client.QueryRow(ctx, sql, chapterID)
	chapter := &pb.GetChapterResponse{}
	err := row.Scan(&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum, &chapter.RegulationID, &chapter.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return chapter, fmt.Errorf("message: %s, code: %s", pgErr.Message, pgErr.Code)
		}
		return chapter, err
	}

	return chapter, nil
}

// GetAll returns all chapters associated with the given ID
func (cs *chapterStorage) GetAll(ctx context.Context, regulationID uint64) ([]*pb.ReadOnlyChapter, error) {
	const sql = `SELECT id,name,num,order_num FROM "chapters" WHERE r_id = $1 ORDER BY order_num`

	var chapters []*pb.ReadOnlyChapter

	rows, err := cs.client.Query(ctx, sql, regulationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chapter := &pb.ReadOnlyChapter{}
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
