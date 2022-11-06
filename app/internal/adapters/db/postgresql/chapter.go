package postgressql

import (
	"context"

	"regulations_service/internal/domain/entity"
	client "regulations_service/pkg/client/postgresql"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient) *chapterStorage {
	return &chapterStorage{client: client}
}

// Create returns the ID of the inserted chapter
func (cs *chapterStorage) CreateOne(ctx context.Context, chapter entity.Chapter) (uint64, error) {
	sql := `INSERT INTO chapters ("name", "num", "order_num","r_id") VALUES ($1,$2,$3,$4) RETURNING "id"`

	row := cs.client.QueryRow(ctx, sql, chapter.Name, chapter.Num, chapter.OrderNum, chapter.RegulationID)

	var chapterID uint64

	err := row.Scan(&chapterID)
	return chapterID, err
}

// GetOneById returns an chapter associated with the given ID
func (cs *chapterStorage) GetOneById(ctx context.Context, chapterID uint64) (entity.Chapter, error) {
	const sql = `SELECT id,name,num,order_num,r_id,updated_at FROM "chapters" WHERE id = $1 ORDER BY order_num`
	row := cs.client.QueryRow(ctx, sql, chapterID)
	var chapter entity.Chapter
	err := row.Scan(&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum, &chapter.RegulationID, &chapter.UpdatedAt)
	if err != nil {
		return chapter, err
	}

	return chapter, nil
}

// GetAllById returns all chapters associated with the given ID
func (cs *chapterStorage) GetAllForRegulation(ctx context.Context, regulationID uint64) ([]entity.Chapter, error) {
	const sql = `SELECT id,name,num,order_num FROM "chapters" WHERE r_id = $1 ORDER BY order_num`

	var chapters []entity.Chapter

	rows, err := cs.client.Query(ctx, sql, regulationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chapter := entity.Chapter{}
		if err = rows.Scan(
			&chapter.ID, &chapter.Name, &chapter.Num, &chapter.OrderNum,
		); err != nil {
			return nil, err
		}

		chapters = append(chapters, chapter)
	}

	return chapters, nil

}

// Delete
func (cs *chapterStorage) DeleteAllForRegulation(ctx context.Context, regulationID uint64) error {
	const sql1 = `delete from chapters where r_id=$1`
	_, err := cs.client.Exec(ctx, sql1, regulationID)
	return err
}
