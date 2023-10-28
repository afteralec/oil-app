// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: player_permissions.sql

package queries

import (
	"context"
	"database/sql"
)

const createPlayerPermission = `-- name: CreatePlayerPermission :execresult
INSERT INTO player_permissions (pid, permission) VALUES (?, ?)
`

type CreatePlayerPermissionParams struct {
	Pid        int64
	Permission string
}

func (q *Queries) CreatePlayerPermission(ctx context.Context, arg CreatePlayerPermissionParams) (sql.Result, error) {
	return q.exec(ctx, q.createPlayerPermissionStmt, createPlayerPermission, arg.Pid, arg.Permission)
}

const listPlayerPermissions = `-- name: ListPlayerPermissions :many
SELECT (id, pid, permission) FROM player_permissions WHERE pid = ?
`

func (q *Queries) ListPlayerPermissions(ctx context.Context, pid int64) ([]interface{}, error) {
	rows, err := q.query(ctx, q.listPlayerPermissionsStmt, listPlayerPermissions, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var column_1 interface{}
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
