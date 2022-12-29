package postgressql

import (
	"context"

	client "read-only_reader_service/pkg/client/postgresql"

	pb "github.com/i-b8o/read-only_contracts/pb/reader/v1"
)

type chapterStorage struct {
	client client.PostgreSQLClient
}

func NewChapterStorage(client client.PostgreSQLClient) *chapterStorage {
	return &chapterStorage{client: client}
}

// Get returns an chapter associated with the given ID
func (cs *chapterStorage) Get(ctx context.Context, chapterID uint64) (*pb.ReaderChapter, error) {
	const sql = `SELECT name,num,order_num,doc_id FROM "chapter" WHERE id = $1 ORDER BY order_num`
	row := cs.client.QueryRow(ctx, sql, chapterID)
	chapter := &pb.ReaderChapter{}
	err := row.Scan(&chapter.Name, &chapter.Num, &chapter.OrderNum, &chapter.DocID)
	if err != nil {
		return chapter, err
	}
	return chapter, nil
}

// GetAll returns all chapters associated with the given ID
func (cs *chapterStorage) GetAll(ctx context.Context, docID uint64) ([]*pb.ReaderChapter, error) {
	const sql = `SELECT id,name,num,order_num FROM "chapter" WHERE doc_id = $1 ORDER BY order_num`

	var chapters []*pb.ReaderChapter

	rows, err := cs.client.Query(ctx, sql, docID)
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
