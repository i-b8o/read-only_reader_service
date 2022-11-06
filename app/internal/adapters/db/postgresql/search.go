package postgressql

import (
	"context"
	"fmt"
	"regulations_service/internal/domain/entity"
	client "regulations_service/pkg/client/postgresql"
)

type searchStorage struct {
	client client.PostgreSQLClient
}

func NewSearchStorage(client client.PostgreSQLClient) *searchStorage {
	return &searchStorage{client: client}
}

func (ss *searchStorage) SearchPargaraphs(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	sql := `SELECT p.id, p.content, c.name, r.name, c.updated_at, count(*) OVER() AS full_count from paragraphs AS p INNER JOIN chapters as c ON c.id = p.c_id INNER JOIN regulations AS r ON c.r_id = r.id  WHERE p.ts @@ phraseto_tsquery('russian',$1)`
	// Pagination
	if len(params) == 3 {
		// sql += fmt.Sprintf(` AND (c.updated_at, p.id) > ('%s' :: TIMESTAMPTZ, '%s') ORDER BY c.updated_at, p.id LIMIT %s`, params[0], params[1], params[2])
		sql += fmt.Sprintf(` OFFSET %s LIMIT %s`, params[0], params[1])
	}
	// else if len(params) == 1 { // First page
	// 	sql += fmt.Sprintf(` LIMIT %s`, params[0])
	// }
	var searchResults []entity.Search
	rows, err := ss.client.Query(ctx, sql, searchQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		search := entity.Search{}
		if err = rows.Scan(
			&search.PID, &search.Text, &search.CName, &search.RName, &search.UpdatedAt, &search.Count,
		); err != nil {
			return nil, err
		}

		searchResults = append(searchResults, search)
	}

	return searchResults, nil
}

func (ss *searchStorage) SearchChapters(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	sql := `SELECT c.id, c.name, r.name, c.updated_at, count(*) OVER() AS full_count from chapters AS c INNER JOIN regulations as r ON c.r_id = r.id WHERE c.ts @@ phraseto_tsquery('russian',$1)`
	// Pagination
	if len(params) == 2 {
		// sql += fmt.Sprintf(` AND (c.updated_at, c.id) > ('%s' :: TIMESTAMPTZ, '%s') ORDER BY c.updated_at, c.id LIMIT %s`, params[0], params[1], params[2])
		sql += fmt.Sprintf(` OFFSET %s LIMIT %s`, params[0], params[1])
	}
	// else if len(params) == 1 { // First page
	// 	sql += fmt.Sprintf(` LIMIT %s`, params[0])
	// }
	var searchResults []entity.Search
	rows, err := ss.client.Query(ctx, sql, searchQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		search := entity.Search{}
		if err = rows.Scan(
			&search.CID, &search.Text, &search.RName, &search.UpdatedAt, &search.Count,
		); err != nil {
			return nil, err
		}

		searchResults = append(searchResults, search)
	}

	return searchResults, nil
}

func (ss *searchStorage) SearchRegulations(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	sql := `SELECT id, name, title, updated_at, count(*) OVER() AS full_count from regulations WHERE ts @@ phraseto_tsquery('russian',$1)`
	// Pagination
	if len(params) == 2 {
		// sql += fmt.Sprintf(` AND (updated_at, id) > ('%s' :: TIMESTAMPTZ, '%s') ORDER BY updated_at, id LIMIT %s`, params[0], params[1], params[2])
		sql += fmt.Sprintf(` OFFSET %s LIMIT %s`, params[0], params[1])
	}
	//  else if len(params) == 1 { // First page
	// 	sql += fmt.Sprintf(` LIMIT %s`, params[0])
	// }
	var searchResults []entity.Search
	rows, err := ss.client.Query(ctx, sql, searchQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		search := entity.Search{}
		if err = rows.Scan(
			&search.RID, &search.RName, &search.Text, &search.UpdatedAt, &search.Count,
		); err != nil {
			return nil, err
		}

		searchResults = append(searchResults, search)
	}

	return searchResults, nil
}

func (ss *searchStorage) Search(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	sql := `SELECT r_id, r_name, c_id, c_name, p_id, text, count(*) OVER() AS full_count FROM reg_search WHERE ts @@ phraseto_tsquery('russian',$1) ORDER BY ts_rank(ts, phraseto_tsquery('russian',$1))`

	var searchResults []entity.Search
	if len(params) == 2 {
		sql += fmt.Sprintf(` OFFSET %s LIMIT %s`, params[0], params[1])
	}

	rows, err := ss.client.Query(ctx, sql, searchQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		search := entity.Search{}
		if err = rows.Scan(
			&search.RID, &search.RName, &search.CID, &search.CName, &search.PID, &search.Text, &search.Count,
		); err != nil {
			return nil, err
		}

		searchResults = append(searchResults, search)
	}

	return searchResults, nil
}

func (ss *searchStorage) SearchLike(ctx context.Context, searchQuery string, params ...string) ([]entity.Search, error) {
	sql := `SELECT r_id, r_name, c_id, c_name, p_id, text, count(*) OVER() AS full_count from reg_search where text like '%` + searchQuery + `%'`
	if len(params) == 2 {
		sql += fmt.Sprintf(` OFFSET %s LIMIT %s`, params[0], params[1])
	}
	var searchResults []entity.Search
	rows, err := ss.client.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		search := entity.Search{}
		if err = rows.Scan(
			&search.RID, &search.RName, &search.CID, &search.CName, &search.PID, &search.Text, &search.Count,
		); err != nil {
			return nil, err
		}

		searchResults = append(searchResults, search)
	}

	return searchResults, nil
}
