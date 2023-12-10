// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: request.sql

package queries

import (
	"context"
	"database/sql"
)

const addCommentToRequest = `-- name: AddCommentToRequest :execresult
INSERT INTO request_comments (text, pid, rid, vid) VALUES (?, ?, ?, ?)
`

type AddCommentToRequestParams struct {
	Text string
	Pid  int64
	Rid  int64
	Vid  int64
}

func (q *Queries) AddCommentToRequest(ctx context.Context, arg AddCommentToRequestParams) (sql.Result, error) {
	return q.exec(ctx, q.addCommentToRequestStmt, addCommentToRequest,
		arg.Text,
		arg.Pid,
		arg.Rid,
		arg.Vid,
	)
}

const addCommentToRequestField = `-- name: AddCommentToRequestField :execresult
INSERT INTO request_comments (text, field, pid, rid, vid) VALUES (?, ?, ?, ?, ?)
`

type AddCommentToRequestFieldParams struct {
	Text  string
	Field string
	Pid   int64
	Rid   int64
	Vid   int64
}

func (q *Queries) AddCommentToRequestField(ctx context.Context, arg AddCommentToRequestFieldParams) (sql.Result, error) {
	return q.exec(ctx, q.addCommentToRequestFieldStmt, addCommentToRequestField,
		arg.Text,
		arg.Field,
		arg.Pid,
		arg.Rid,
		arg.Vid,
	)
}

const addReplyToComment = `-- name: AddReplyToComment :execresult
INSERT INTO request_comments (text, cid, pid, rid, vid) VALUES (?, ?, ?, ?, ?)
`

type AddReplyToCommentParams struct {
	Text string
	Cid  int64
	Pid  int64
	Rid  int64
	Vid  int64
}

func (q *Queries) AddReplyToComment(ctx context.Context, arg AddReplyToCommentParams) (sql.Result, error) {
	return q.exec(ctx, q.addReplyToCommentStmt, addReplyToComment,
		arg.Text,
		arg.Cid,
		arg.Pid,
		arg.Rid,
		arg.Vid,
	)
}

const addReplyToFieldComment = `-- name: AddReplyToFieldComment :execresult
INSERT INTO request_comments (text, field, cid, pid, rid, vid) VALUES (?, ?, ?, ?, ?, ?)
`

type AddReplyToFieldCommentParams struct {
	Text  string
	Field string
	Cid   int64
	Pid   int64
	Rid   int64
	Vid   int64
}

func (q *Queries) AddReplyToFieldComment(ctx context.Context, arg AddReplyToFieldCommentParams) (sql.Result, error) {
	return q.exec(ctx, q.addReplyToFieldCommentStmt, addReplyToFieldComment,
		arg.Text,
		arg.Field,
		arg.Cid,
		arg.Pid,
		arg.Rid,
		arg.Vid,
	)
}

const countOpenRequests = `-- name: CountOpenRequests :one
SELECT
  COUNT(*)
FROM
  requests
WHERE
  pid = ? AND status != "Archived" AND status != "Canceled"
`

func (q *Queries) CountOpenRequests(ctx context.Context, pid int64) (int64, error) {
	row := q.queryRow(ctx, q.countOpenRequestsStmt, countOpenRequests, pid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createRequest = `-- name: CreateRequest :execresult
INSERT INTO requests (type, pid) VALUES (?, ?)
`

type CreateRequestParams struct {
	Type string
	Pid  int64
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) (sql.Result, error) {
	return q.exec(ctx, q.createRequestStmt, createRequest, arg.Type, arg.Pid)
}

const getRequest = `-- name: GetRequest :one
SELECT type, status, created_at, updated_at, vid, pid, id, new FROM requests WHERE id = ?
`

func (q *Queries) GetRequest(ctx context.Context, id int64) (Request, error) {
	row := q.queryRow(ctx, q.getRequestStmt, getRequest, id)
	var i Request
	err := row.Scan(
		&i.Type,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Vid,
		&i.Pid,
		&i.ID,
		&i.New,
	)
	return i, err
}

const getRequestComment = `-- name: GetRequestComment :one
SELECT created_at, updated_at, deleted_at, text, field, deleted, cid, rid, vid, pid, id FROM request_comments WHERE id = ?
`

func (q *Queries) GetRequestComment(ctx context.Context, id int64) (RequestComment, error) {
	row := q.queryRow(ctx, q.getRequestCommentStmt, getRequestComment, id)
	var i RequestComment
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Text,
		&i.Field,
		&i.Deleted,
		&i.Cid,
		&i.Rid,
		&i.Vid,
		&i.Pid,
		&i.ID,
	)
	return i, err
}

const listCommentsForRequest = `-- name: ListCommentsForRequest :many
SELECT created_at, updated_at, deleted_at, text, field, deleted, cid, rid, vid, pid, id FROM request_comments WHERE rid = ?
`

func (q *Queries) ListCommentsForRequest(ctx context.Context, rid int64) ([]RequestComment, error) {
	rows, err := q.query(ctx, q.listCommentsForRequestStmt, listCommentsForRequest, rid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RequestComment
	for rows.Next() {
		var i RequestComment
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Text,
			&i.Field,
			&i.Deleted,
			&i.Cid,
			&i.Rid,
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

const listRepliesToComment = `-- name: ListRepliesToComment :many
SELECT created_at, updated_at, deleted_at, text, field, deleted, cid, rid, vid, pid, id FROM request_comments WHERE cid = ?
`

func (q *Queries) ListRepliesToComment(ctx context.Context, cid int64) ([]RequestComment, error) {
	rows, err := q.query(ctx, q.listRepliesToCommentStmt, listRepliesToComment, cid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RequestComment
	for rows.Next() {
		var i RequestComment
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Text,
			&i.Field,
			&i.Deleted,
			&i.Cid,
			&i.Rid,
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
SELECT type, status, created_at, updated_at, vid, pid, id, new FROM requests WHERE pid = ?
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
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Vid,
			&i.Pid,
			&i.ID,
			&i.New,
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
