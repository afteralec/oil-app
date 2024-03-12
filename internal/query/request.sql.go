// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: request.sql

package query

import (
	"context"
	"database/sql"
)

const countOpenCharacterApplicationsForPlayer = `-- name: CountOpenCharacterApplicationsForPlayer :one
SELECT
  COUNT(*)
FROM
  requests
WHERE
  pid = ?
AND
  type = "CharacterApplication"
AND
  status != "Archived"
AND
  status != "Canceled"
`

func (q *Queries) CountOpenCharacterApplicationsForPlayer(ctx context.Context, pid int64) (int64, error) {
	row := q.queryRow(ctx, q.countOpenCharacterApplicationsForPlayerStmt, countOpenCharacterApplicationsForPlayer, pid)
	var count int64
	err := row.Scan(&count)
	return count, err
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

const countUnresolvedCommentsForRequest = `-- name: CountUnresolvedCommentsForRequest :one
SELECT COUNT(*) FROM request_comments WHERE rid = ? AND resolved = false
`

func (q *Queries) CountUnresolvedCommentsForRequest(ctx context.Context, rid int64) (int64, error) {
	row := q.queryRow(ctx, q.countUnresolvedCommentsForRequestStmt, countUnresolvedCommentsForRequest, rid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countUnresolvedCommentsForRequestField = `-- name: CountUnresolvedCommentsForRequestField :one
SELECT COUNT(*) FROM request_comments WHERE rid = ? AND field = ? AND resolved = false
`

type CountUnresolvedCommentsForRequestFieldParams struct {
	RID   int64
	Field string
}

func (q *Queries) CountUnresolvedCommentsForRequestField(ctx context.Context, arg CountUnresolvedCommentsForRequestFieldParams) (int64, error) {
	row := q.queryRow(ctx, q.countUnresolvedCommentsForRequestFieldStmt, countUnresolvedCommentsForRequestField, arg.RID, arg.Field)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createCharacterApplicationContent = `-- name: CreateCharacterApplicationContent :exec
INSERT INTO
  character_application_content 
  (gender, name, short_description, description, backstory, rid) 
VALUES 
  ("", "", "", "", "", ?)
`

func (q *Queries) CreateCharacterApplicationContent(ctx context.Context, rid int64) error {
	_, err := q.exec(ctx, q.createCharacterApplicationContentStmt, createCharacterApplicationContent, rid)
	return err
}

const createCharacterApplicationContentReview = `-- name: CreateCharacterApplicationContentReview :exec
INSERT INTO
  character_application_content_review
  (gender, name, short_description, description, backstory, rid) 
VALUES 
  (?, ?, ?, ?, ?, ?)
`

type CreateCharacterApplicationContentReviewParams struct {
	Gender           string `json:"gender"`
	Name             string `json:"name"`
	ShortDescription string `json:"sdesc"`
	Description      string `json:"desc"`
	Backstory        string `json:"backstory"`
	RID              int64  `json:"-"`
}

func (q *Queries) CreateCharacterApplicationContentReview(ctx context.Context, arg CreateCharacterApplicationContentReviewParams) error {
	_, err := q.exec(ctx, q.createCharacterApplicationContentReviewStmt, createCharacterApplicationContentReview,
		arg.Gender,
		arg.Name,
		arg.ShortDescription,
		arg.Description,
		arg.Backstory,
		arg.RID,
	)
	return err
}

const createHistoryForCharacterApplication = `-- name: CreateHistoryForCharacterApplication :exec
INSERT INTO
  character_application_content_history
  (gender, name, short_description, description, backstory, rid, vid)
SELECT 
  gender, name, short_description, description, backstory, rid, requests.vid
FROM
  character_application_content
JOIN
  requests
ON
  requests.id = character_application_content.rid
WHERE
  character_application_content.rid = ?
`

func (q *Queries) CreateHistoryForCharacterApplication(ctx context.Context, rid int64) error {
	_, err := q.exec(ctx, q.createHistoryForCharacterApplicationStmt, createHistoryForCharacterApplication, rid)
	return err
}

const createHistoryForRequestStatusChange = `-- name: CreateHistoryForRequestStatusChange :exec
INSERT INTO 
  request_status_change_history
  (rid, vid, status, pid)
VALUES
  (?, (SELECT vid FROM requests WHERE requests.id = rid), (SELECT status FROM requests WHERE requests.id = rid), ?)
`

type CreateHistoryForRequestStatusChangeParams struct {
	RID int64
	PID int64
}

func (q *Queries) CreateHistoryForRequestStatusChange(ctx context.Context, arg CreateHistoryForRequestStatusChangeParams) error {
	_, err := q.exec(ctx, q.createHistoryForRequestStatusChangeStmt, createHistoryForRequestStatusChange, arg.RID, arg.PID)
	return err
}

const createRequest = `-- name: CreateRequest :execresult
INSERT INTO requests (type, pid) VALUES (?, ?)
`

type CreateRequestParams struct {
	Type string
	PID  int64
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) (sql.Result, error) {
	return q.exec(ctx, q.createRequestStmt, createRequest, arg.Type, arg.PID)
}

const createRequestChangeRequest = `-- name: CreateRequestChangeRequest :exec
INSERT INTO request_change_requests (rid, field, pid, text) VALUES (?, ?, ?, ?)
`

type CreateRequestChangeRequestParams struct {
	RID   int64
	Field string
	PID   int64
	Text  string
}

func (q *Queries) CreateRequestChangeRequest(ctx context.Context, arg CreateRequestChangeRequestParams) error {
	_, err := q.exec(ctx, q.createRequestChangeRequestStmt, createRequestChangeRequest,
		arg.RID,
		arg.Field,
		arg.PID,
		arg.Text,
	)
	return err
}

const createRequestComment = `-- name: CreateRequestComment :execresult
INSERT INTO
  request_comments (text, field, pid, rid, vid) 
VALUES
  (?, ?, ?, ?, (SELECT vid FROM requests WHERE requests.id = rid))
`

type CreateRequestCommentParams struct {
	Text  string
	Field string
	PID   int64
	RID   int64
}

func (q *Queries) CreateRequestComment(ctx context.Context, arg CreateRequestCommentParams) (sql.Result, error) {
	return q.exec(ctx, q.createRequestCommentStmt, createRequestComment,
		arg.Text,
		arg.Field,
		arg.PID,
		arg.RID,
	)
}

const getCharacterApplication = `-- name: GetCharacterApplication :one
SELECT
  character_application_content.created_at, character_application_content.updated_at, character_application_content.backstory, character_application_content.description, character_application_content.short_description, character_application_content.name, character_application_content.gender, character_application_content.rid, character_application_content.id, requests.created_at, requests.updated_at, requests.type, requests.status, requests.rpid, requests.pid, requests.id, requests.vid
FROM
  requests
JOIN
  character_application_content
ON
  character_application_content.rid = requests.id
WHERE
  requests.id = ?
`

type GetCharacterApplicationRow struct {
	CharacterApplicationContent CharacterApplicationContent
	Request                     Request
}

func (q *Queries) GetCharacterApplication(ctx context.Context, id int64) (GetCharacterApplicationRow, error) {
	row := q.queryRow(ctx, q.getCharacterApplicationStmt, getCharacterApplication, id)
	var i GetCharacterApplicationRow
	err := row.Scan(
		&i.CharacterApplicationContent.CreatedAt,
		&i.CharacterApplicationContent.UpdatedAt,
		&i.CharacterApplicationContent.Backstory,
		&i.CharacterApplicationContent.Description,
		&i.CharacterApplicationContent.ShortDescription,
		&i.CharacterApplicationContent.Name,
		&i.CharacterApplicationContent.Gender,
		&i.CharacterApplicationContent.RID,
		&i.CharacterApplicationContent.ID,
		&i.Request.CreatedAt,
		&i.Request.UpdatedAt,
		&i.Request.Type,
		&i.Request.Status,
		&i.Request.RPID,
		&i.Request.PID,
		&i.Request.ID,
		&i.Request.VID,
	)
	return i, err
}

const getCharacterApplicationContent = `-- name: GetCharacterApplicationContent :one
SELECT created_at, updated_at, backstory, description, short_description, name, gender, rid, id FROM character_application_content WHERE id = ?
`

func (q *Queries) GetCharacterApplicationContent(ctx context.Context, id int64) (CharacterApplicationContent, error) {
	row := q.queryRow(ctx, q.getCharacterApplicationContentStmt, getCharacterApplicationContent, id)
	var i CharacterApplicationContent
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Backstory,
		&i.Description,
		&i.ShortDescription,
		&i.Name,
		&i.Gender,
		&i.RID,
		&i.ID,
	)
	return i, err
}

const getCharacterApplicationContentForRequest = `-- name: GetCharacterApplicationContentForRequest :one
SELECT created_at, updated_at, backstory, description, short_description, name, gender, rid, id FROM character_application_content WHERE rid = ?
`

func (q *Queries) GetCharacterApplicationContentForRequest(ctx context.Context, rid int64) (CharacterApplicationContent, error) {
	row := q.queryRow(ctx, q.getCharacterApplicationContentForRequestStmt, getCharacterApplicationContentForRequest, rid)
	var i CharacterApplicationContent
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Backstory,
		&i.Description,
		&i.ShortDescription,
		&i.Name,
		&i.Gender,
		&i.RID,
		&i.ID,
	)
	return i, err
}

const getCharacterApplicationContentReviewForRequest = `-- name: GetCharacterApplicationContentReviewForRequest :one
SELECT created_at, updated_at, name, gender, short_description, description, backstory, rid, id FROM character_application_content_review WHERE rid = ?
`

func (q *Queries) GetCharacterApplicationContentReviewForRequest(ctx context.Context, rid int64) (CharacterApplicationContentReview, error) {
	row := q.queryRow(ctx, q.getCharacterApplicationContentReviewForRequestStmt, getCharacterApplicationContentReviewForRequest, rid)
	var i CharacterApplicationContentReview
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Gender,
		&i.ShortDescription,
		&i.Description,
		&i.Backstory,
		&i.RID,
		&i.ID,
	)
	return i, err
}

const getCommentWithAuthor = `-- name: GetCommentWithAuthor :one
SELECT 
  players.created_at, players.updated_at, players.pw_hash, players.username, players.id, request_comments.created_at, request_comments.updated_at, request_comments.deleted_at, request_comments.text, request_comments.field, request_comments.rid, request_comments.pid, request_comments.cid, request_comments.id, request_comments.vid, request_comments.deleted, request_comments.resolved
FROM 
  request_comments 
JOIN
  players
ON
  request_comments.pid = players.id
WHERE 
  request_comments.id = ?
`

type GetCommentWithAuthorRow struct {
	Player         Player
	RequestComment RequestComment
}

func (q *Queries) GetCommentWithAuthor(ctx context.Context, id int64) (GetCommentWithAuthorRow, error) {
	row := q.queryRow(ctx, q.getCommentWithAuthorStmt, getCommentWithAuthor, id)
	var i GetCommentWithAuthorRow
	err := row.Scan(
		&i.Player.CreatedAt,
		&i.Player.UpdatedAt,
		&i.Player.PwHash,
		&i.Player.Username,
		&i.Player.ID,
		&i.RequestComment.CreatedAt,
		&i.RequestComment.UpdatedAt,
		&i.RequestComment.DeletedAt,
		&i.RequestComment.Text,
		&i.RequestComment.Field,
		&i.RequestComment.RID,
		&i.RequestComment.PID,
		&i.RequestComment.CID,
		&i.RequestComment.ID,
		&i.RequestComment.VID,
		&i.RequestComment.Deleted,
		&i.RequestComment.Resolved,
	)
	return i, err
}

const getCurrentRequestChangeRequestForRequestField = `-- name: GetCurrentRequestChangeRequestForRequestField :one
SELECT created_at, updated_at, text, field, rid, pid, id, old FROM request_change_requests WHERE rid = ? AND field = ? AND old = false
`

type GetCurrentRequestChangeRequestForRequestFieldParams struct {
	RID   int64
	Field string
}

func (q *Queries) GetCurrentRequestChangeRequestForRequestField(ctx context.Context, arg GetCurrentRequestChangeRequestForRequestFieldParams) (RequestChangeRequest, error) {
	row := q.queryRow(ctx, q.getCurrentRequestChangeRequestForRequestFieldStmt, getCurrentRequestChangeRequestForRequestField, arg.RID, arg.Field)
	var i RequestChangeRequest
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Text,
		&i.Field,
		&i.RID,
		&i.PID,
		&i.ID,
		&i.Old,
	)
	return i, err
}

const getRequest = `-- name: GetRequest :one
SELECT created_at, updated_at, type, status, rpid, pid, id, vid FROM requests WHERE id = ?
`

func (q *Queries) GetRequest(ctx context.Context, id int64) (Request, error) {
	row := q.queryRow(ctx, q.getRequestStmt, getRequest, id)
	var i Request
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.RPID,
		&i.PID,
		&i.ID,
		&i.VID,
	)
	return i, err
}

const incrementRequestVersion = `-- name: IncrementRequestVersion :exec
UPDATE requests SET vid = vid + 1 WHERE id = ?
`

func (q *Queries) IncrementRequestVersion(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.incrementRequestVersionStmt, incrementRequestVersion, id)
	return err
}

const listCharacterApplicationContentForPlayer = `-- name: ListCharacterApplicationContentForPlayer :many
SELECT
  created_at, updated_at, backstory, description, short_description, name, gender, rid, id
FROM
  character_application_content 
WHERE
  rid
IN (SELECT id FROM requests WHERE pid = ?)
`

func (q *Queries) ListCharacterApplicationContentForPlayer(ctx context.Context, pid int64) ([]CharacterApplicationContent, error) {
	rows, err := q.query(ctx, q.listCharacterApplicationContentForPlayerStmt, listCharacterApplicationContentForPlayer, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CharacterApplicationContent
	for rows.Next() {
		var i CharacterApplicationContent
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Backstory,
			&i.Description,
			&i.ShortDescription,
			&i.Name,
			&i.Gender,
			&i.RID,
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

const listCharacterApplicationsForPlayer = `-- name: ListCharacterApplicationsForPlayer :many
SELECT
  character_application_content.created_at, character_application_content.updated_at, character_application_content.backstory, character_application_content.description, character_application_content.short_description, character_application_content.name, character_application_content.gender, character_application_content.rid, character_application_content.id, players.created_at, players.updated_at, players.pw_hash, players.username, players.id, requests.created_at, requests.updated_at, requests.type, requests.status, requests.rpid, requests.pid, requests.id, requests.vid
FROM
  requests
JOIN
  character_application_content
ON
  requests.id = character_application_content.rid
JOIN
  players
ON
  players.id = requests.pid
WHERE
  requests.pid = ?
AND
  requests.type = "CharacterApplication"
AND
  requests.status != "Archived"
AND
  requests.status != "Canceled"
`

type ListCharacterApplicationsForPlayerRow struct {
	CharacterApplicationContent CharacterApplicationContent
	Player                      Player
	Request                     Request
}

func (q *Queries) ListCharacterApplicationsForPlayer(ctx context.Context, pid int64) ([]ListCharacterApplicationsForPlayerRow, error) {
	rows, err := q.query(ctx, q.listCharacterApplicationsForPlayerStmt, listCharacterApplicationsForPlayer, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCharacterApplicationsForPlayerRow
	for rows.Next() {
		var i ListCharacterApplicationsForPlayerRow
		if err := rows.Scan(
			&i.CharacterApplicationContent.CreatedAt,
			&i.CharacterApplicationContent.UpdatedAt,
			&i.CharacterApplicationContent.Backstory,
			&i.CharacterApplicationContent.Description,
			&i.CharacterApplicationContent.ShortDescription,
			&i.CharacterApplicationContent.Name,
			&i.CharacterApplicationContent.Gender,
			&i.CharacterApplicationContent.RID,
			&i.CharacterApplicationContent.ID,
			&i.Player.CreatedAt,
			&i.Player.UpdatedAt,
			&i.Player.PwHash,
			&i.Player.Username,
			&i.Player.ID,
			&i.Request.CreatedAt,
			&i.Request.UpdatedAt,
			&i.Request.Type,
			&i.Request.Status,
			&i.Request.RPID,
			&i.Request.PID,
			&i.Request.ID,
			&i.Request.VID,
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

const listCommentsForRequest = `-- name: ListCommentsForRequest :many
SELECT created_at, updated_at, deleted_at, text, field, rid, pid, cid, id, vid, deleted, resolved FROM request_comments WHERE rid = ?
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
			&i.RID,
			&i.PID,
			&i.CID,
			&i.ID,
			&i.VID,
			&i.Deleted,
			&i.Resolved,
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

const listCommentsForRequestFieldWithAuthor = `-- name: ListCommentsForRequestFieldWithAuthor :many
SELECT
  players.created_at, players.updated_at, players.pw_hash, players.username, players.id, request_comments.created_at, request_comments.updated_at, request_comments.deleted_at, request_comments.text, request_comments.field, request_comments.rid, request_comments.pid, request_comments.cid, request_comments.id, request_comments.vid, request_comments.deleted, request_comments.resolved
FROM
  request_comments
JOIN
  players
ON
  request_comments.pid = players.id
WHERE
  field = ? AND rid = ?
`

type ListCommentsForRequestFieldWithAuthorParams struct {
	Field string
	RID   int64
}

type ListCommentsForRequestFieldWithAuthorRow struct {
	Player         Player
	RequestComment RequestComment
}

func (q *Queries) ListCommentsForRequestFieldWithAuthor(ctx context.Context, arg ListCommentsForRequestFieldWithAuthorParams) ([]ListCommentsForRequestFieldWithAuthorRow, error) {
	rows, err := q.query(ctx, q.listCommentsForRequestFieldWithAuthorStmt, listCommentsForRequestFieldWithAuthor, arg.Field, arg.RID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCommentsForRequestFieldWithAuthorRow
	for rows.Next() {
		var i ListCommentsForRequestFieldWithAuthorRow
		if err := rows.Scan(
			&i.Player.CreatedAt,
			&i.Player.UpdatedAt,
			&i.Player.PwHash,
			&i.Player.Username,
			&i.Player.ID,
			&i.RequestComment.CreatedAt,
			&i.RequestComment.UpdatedAt,
			&i.RequestComment.DeletedAt,
			&i.RequestComment.Text,
			&i.RequestComment.Field,
			&i.RequestComment.RID,
			&i.RequestComment.PID,
			&i.RequestComment.CID,
			&i.RequestComment.ID,
			&i.RequestComment.VID,
			&i.RequestComment.Deleted,
			&i.RequestComment.Resolved,
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

const listCommentsForRequestWithAuthor = `-- name: ListCommentsForRequestWithAuthor :many
SELECT
  players.created_at, players.updated_at, players.pw_hash, players.username, players.id, request_comments.created_at, request_comments.updated_at, request_comments.deleted_at, request_comments.text, request_comments.field, request_comments.rid, request_comments.pid, request_comments.cid, request_comments.id, request_comments.vid, request_comments.deleted, request_comments.resolved
FROM
  request_comments
JOIN
  players
ON
  request_comments.pid = players.id
WHERE
  rid = ?
`

type ListCommentsForRequestWithAuthorRow struct {
	Player         Player
	RequestComment RequestComment
}

func (q *Queries) ListCommentsForRequestWithAuthor(ctx context.Context, rid int64) ([]ListCommentsForRequestWithAuthorRow, error) {
	rows, err := q.query(ctx, q.listCommentsForRequestWithAuthorStmt, listCommentsForRequestWithAuthor, rid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCommentsForRequestWithAuthorRow
	for rows.Next() {
		var i ListCommentsForRequestWithAuthorRow
		if err := rows.Scan(
			&i.Player.CreatedAt,
			&i.Player.UpdatedAt,
			&i.Player.PwHash,
			&i.Player.Username,
			&i.Player.ID,
			&i.RequestComment.CreatedAt,
			&i.RequestComment.UpdatedAt,
			&i.RequestComment.DeletedAt,
			&i.RequestComment.Text,
			&i.RequestComment.Field,
			&i.RequestComment.RID,
			&i.RequestComment.PID,
			&i.RequestComment.CID,
			&i.RequestComment.ID,
			&i.RequestComment.VID,
			&i.RequestComment.Deleted,
			&i.RequestComment.Resolved,
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

const listOpenCharacterApplications = `-- name: ListOpenCharacterApplications :many
SELECT 
  character_application_content.created_at, character_application_content.updated_at, character_application_content.backstory, character_application_content.description, character_application_content.short_description, character_application_content.name, character_application_content.gender, character_application_content.rid, character_application_content.id, players.created_at, players.updated_at, players.pw_hash, players.username, players.id, requests.created_at, requests.updated_at, requests.type, requests.status, requests.rpid, requests.pid, requests.id, requests.vid
FROM 
  requests
JOIN 
  character_application_content
ON 
  requests.id = character_application_content.rid
JOIN 
  players
ON 
  players.id = requests.pid
WHERE 
  requests.type = "CharacterApplication"
AND 
  requests.status = "Submitted"
OR 
  requests.status = "InReview"
OR 
  requests.status = "Reviewed"
`

type ListOpenCharacterApplicationsRow struct {
	CharacterApplicationContent CharacterApplicationContent
	Player                      Player
	Request                     Request
}

func (q *Queries) ListOpenCharacterApplications(ctx context.Context) ([]ListOpenCharacterApplicationsRow, error) {
	rows, err := q.query(ctx, q.listOpenCharacterApplicationsStmt, listOpenCharacterApplications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListOpenCharacterApplicationsRow
	for rows.Next() {
		var i ListOpenCharacterApplicationsRow
		if err := rows.Scan(
			&i.CharacterApplicationContent.CreatedAt,
			&i.CharacterApplicationContent.UpdatedAt,
			&i.CharacterApplicationContent.Backstory,
			&i.CharacterApplicationContent.Description,
			&i.CharacterApplicationContent.ShortDescription,
			&i.CharacterApplicationContent.Name,
			&i.CharacterApplicationContent.Gender,
			&i.CharacterApplicationContent.RID,
			&i.CharacterApplicationContent.ID,
			&i.Player.CreatedAt,
			&i.Player.UpdatedAt,
			&i.Player.PwHash,
			&i.Player.Username,
			&i.Player.ID,
			&i.Request.CreatedAt,
			&i.Request.UpdatedAt,
			&i.Request.Type,
			&i.Request.Status,
			&i.Request.RPID,
			&i.Request.PID,
			&i.Request.ID,
			&i.Request.VID,
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
SELECT created_at, updated_at, type, status, rpid, pid, id, vid FROM requests WHERE pid = ?
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
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Type,
			&i.Status,
			&i.RPID,
			&i.PID,
			&i.ID,
			&i.VID,
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

const updateCharacterApplicationContentBackstory = `-- name: UpdateCharacterApplicationContentBackstory :exec
UPDATE character_application_content SET backstory = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentBackstoryParams struct {
	Backstory string `json:"backstory"`
	RID       int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentBackstory(ctx context.Context, arg UpdateCharacterApplicationContentBackstoryParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentBackstoryStmt, updateCharacterApplicationContentBackstory, arg.Backstory, arg.RID)
	return err
}

const updateCharacterApplicationContentDescription = `-- name: UpdateCharacterApplicationContentDescription :exec
UPDATE character_application_content SET description = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentDescriptionParams struct {
	Description string `json:"desc"`
	RID         int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentDescription(ctx context.Context, arg UpdateCharacterApplicationContentDescriptionParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentDescriptionStmt, updateCharacterApplicationContentDescription, arg.Description, arg.RID)
	return err
}

const updateCharacterApplicationContentGender = `-- name: UpdateCharacterApplicationContentGender :exec
UPDATE character_application_content SET gender = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentGenderParams struct {
	Gender string `json:"gender"`
	RID    int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentGender(ctx context.Context, arg UpdateCharacterApplicationContentGenderParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentGenderStmt, updateCharacterApplicationContentGender, arg.Gender, arg.RID)
	return err
}

const updateCharacterApplicationContentName = `-- name: UpdateCharacterApplicationContentName :exec
UPDATE character_application_content SET name = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentNameParams struct {
	Name string `json:"name"`
	RID  int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentName(ctx context.Context, arg UpdateCharacterApplicationContentNameParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentNameStmt, updateCharacterApplicationContentName, arg.Name, arg.RID)
	return err
}

const updateCharacterApplicationContentReviewBackstory = `-- name: UpdateCharacterApplicationContentReviewBackstory :exec
UPDATE character_application_content_review SET backstory = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentReviewBackstoryParams struct {
	Backstory string `json:"backstory"`
	RID       int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentReviewBackstory(ctx context.Context, arg UpdateCharacterApplicationContentReviewBackstoryParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentReviewBackstoryStmt, updateCharacterApplicationContentReviewBackstory, arg.Backstory, arg.RID)
	return err
}

const updateCharacterApplicationContentReviewDescription = `-- name: UpdateCharacterApplicationContentReviewDescription :exec
UPDATE character_application_content_review SET description = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentReviewDescriptionParams struct {
	Description string `json:"desc"`
	RID         int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentReviewDescription(ctx context.Context, arg UpdateCharacterApplicationContentReviewDescriptionParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentReviewDescriptionStmt, updateCharacterApplicationContentReviewDescription, arg.Description, arg.RID)
	return err
}

const updateCharacterApplicationContentReviewGender = `-- name: UpdateCharacterApplicationContentReviewGender :exec
UPDATE character_application_content_review SET gender = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentReviewGenderParams struct {
	Gender string `json:"gender"`
	RID    int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentReviewGender(ctx context.Context, arg UpdateCharacterApplicationContentReviewGenderParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentReviewGenderStmt, updateCharacterApplicationContentReviewGender, arg.Gender, arg.RID)
	return err
}

const updateCharacterApplicationContentReviewName = `-- name: UpdateCharacterApplicationContentReviewName :exec
UPDATE character_application_content_review SET name = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentReviewNameParams struct {
	Name string `json:"name"`
	RID  int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentReviewName(ctx context.Context, arg UpdateCharacterApplicationContentReviewNameParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentReviewNameStmt, updateCharacterApplicationContentReviewName, arg.Name, arg.RID)
	return err
}

const updateCharacterApplicationContentReviewShortDescription = `-- name: UpdateCharacterApplicationContentReviewShortDescription :exec
UPDATE character_application_content_review SET short_description = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentReviewShortDescriptionParams struct {
	ShortDescription string `json:"sdesc"`
	RID              int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentReviewShortDescription(ctx context.Context, arg UpdateCharacterApplicationContentReviewShortDescriptionParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentReviewShortDescriptionStmt, updateCharacterApplicationContentReviewShortDescription, arg.ShortDescription, arg.RID)
	return err
}

const updateCharacterApplicationContentShortDescription = `-- name: UpdateCharacterApplicationContentShortDescription :exec
UPDATE character_application_content SET short_description = ? WHERE rid = ?
`

type UpdateCharacterApplicationContentShortDescriptionParams struct {
	ShortDescription string `json:"sdesc"`
	RID              int64  `json:"-"`
}

func (q *Queries) UpdateCharacterApplicationContentShortDescription(ctx context.Context, arg UpdateCharacterApplicationContentShortDescriptionParams) error {
	_, err := q.exec(ctx, q.updateCharacterApplicationContentShortDescriptionStmt, updateCharacterApplicationContentShortDescription, arg.ShortDescription, arg.RID)
	return err
}

const updateRequestReviewer = `-- name: UpdateRequestReviewer :exec
UPDATE requests SET rpid = ? WHERE id = ?
`

type UpdateRequestReviewerParams struct {
	RPID int64
	ID   int64
}

func (q *Queries) UpdateRequestReviewer(ctx context.Context, arg UpdateRequestReviewerParams) error {
	_, err := q.exec(ctx, q.updateRequestReviewerStmt, updateRequestReviewer, arg.RPID, arg.ID)
	return err
}

const updateRequestStatus = `-- name: UpdateRequestStatus :exec
UPDATE requests SET status = ? WHERE id = ?
`

type UpdateRequestStatusParams struct {
	Status string
	ID     int64
}

func (q *Queries) UpdateRequestStatus(ctx context.Context, arg UpdateRequestStatusParams) error {
	_, err := q.exec(ctx, q.updateRequestStatusStmt, updateRequestStatus, arg.Status, arg.ID)
	return err
}
