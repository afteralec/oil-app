// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: request.sql

package queries

import (
	"context"
)

const getRequest = `-- name: GetRequest :one
SELECT type, created_at, updated_at, vid, pid, id FROM requests WHERE id = ?
`

func (q *Queries) GetRequest(ctx context.Context, id int64) (Request, error) {
	row := q.queryRow(ctx, q.getRequestStmt, getRequest, id)
	var i Request
	err := row.Scan(
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Vid,
		&i.Pid,
		&i.ID,
	)
	return i, err
}

const listCharacterApplicationsForPlayer = `-- name: ListCharacterApplicationsForPlayer :many
SELECT type, created_at, updated_at, vid, pid, id FROM requests WHERE pid = ? AND type = 'CharacterApplication'
`

func (q *Queries) ListCharacterApplicationsForPlayer(ctx context.Context, pid int64) ([]Request, error) {
	rows, err := q.query(ctx, q.listCharacterApplicationsForPlayerStmt, listCharacterApplicationsForPlayer, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Request
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Vid,
			&i.Pid,
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRequestsForPlayer = `-- name: ListRequestsForPlayer :many
SELECT type, created_at, updated_at, vid, pid, id FROM requests WHERE pid = ?
`

func (q *Queries) ListRequestsForPlayer(ctx context.Context, pid int64) ([]Request, error) {
	rows, err := q.query(ctx, q.listRequestsForPlayerStmt, listRequestsForPlayer, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Request
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Vid,
			&i.Pid,
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
