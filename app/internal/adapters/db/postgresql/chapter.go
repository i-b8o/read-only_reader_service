package postgressql

import (
	"context"

	pb "regulations_read_only_service/internal/pb"
	client "regulations_read_only_service/pkg/client/postgresql"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient) *chapterStorage {
	return &chapterStorage{client: client}
}

// Get returns an chapter associated with the given ID
func (cs *chapterStorage) Get(ctx context.Context, chapterID uint64) (*pb.Chapter, error) {
	const sql = `SELECT id,name,num,order_num,r_id,updated_at FROM "chapters" WHERE id = $1 ORDER BY order_num`
	row := cs.client.QueryRow(ctx, sql, chapterID)
	chapter := &pb.Chapter{}
	err := row.Scan(&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum, &chapter.RegulationID, &chapter.UpdatedAt)
	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

// GetAll returns all chapters associated with the given ID
func (cs *chapterStorage) GetAll(ctx context.Context, regulationID uint64) ([]*pb.Chapter, error) {
	const sql = `SELECT id,name,num,order_num FROM "chapters" WHERE r_id = $1 ORDER BY order_num`

	var chapters []*pb.Chapter

	rows, err := cs.client.Query(ctx, sql, regulationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chapter := &pb.Chapter{}
		if err = rows.Scan(
			&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum,
		); err != nil {
			return nil, err
		}

		chapters = append(chapters, chapter)
	}

	return chapters, nil

}
