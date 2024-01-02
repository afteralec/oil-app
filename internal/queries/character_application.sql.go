// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: character_application.sql

package queries

import (
	"context"
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
