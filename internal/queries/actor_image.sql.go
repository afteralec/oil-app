// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: actor_image.sql

package queries

import (
	"context"
	"database/sql"
)

const createActorImage = `-- name: CreateActorImage :execresult
INSERT INTO actor_images (gender, name, short_description, description) VALUES (?, ?, ?, ?)
`

type CreateActorImageParams struct {
	Gender           string
	Name             string
	ShortDescription string
	Description      string
}

func (q *Queries) CreateActorImage(ctx context.Context, arg CreateActorImageParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageStmt, createActorImage,
		arg.Gender,
		arg.Name,
		arg.ShortDescription,
		arg.Description,
	)
}

const createActorImageCan = `-- name: CreateActorImageCan :execresult
INSERT INTO actor_images_can (can, aiid) VALUES (?, ?)
`

type CreateActorImageCanParams struct {
	Can  string
	AIID int64
}

func (q *Queries) CreateActorImageCan(ctx context.Context, arg CreateActorImageCanParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageCanStmt, createActorImageCan, arg.Can, arg.AIID)
}

const createActorImageCanBe = `-- name: CreateActorImageCanBe :execresult
INSERT INTO actor_images_can_be (can_be, aiid) VALUES (?, ?)
`

type CreateActorImageCanBeParams struct {
	CanBe string
	AIID  int64
}

func (q *Queries) CreateActorImageCanBe(ctx context.Context, arg CreateActorImageCanBeParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageCanBeStmt, createActorImageCanBe, arg.CanBe, arg.AIID)
}

const createActorImageContainerProperties = `-- name: CreateActorImageContainerProperties :execresult
INSERT INTO actor_images_container_properties (aiid, is_container, is_surface_container, liquid_capacity) VALUES (?, ?, ?, ?)
`

type CreateActorImageContainerPropertiesParams struct {
	AIID               int64
	IsContainer        bool
	IsSurfaceContainer bool
	LiquidCapacity     int32
}

func (q *Queries) CreateActorImageContainerProperties(ctx context.Context, arg CreateActorImageContainerPropertiesParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageContainerPropertiesStmt, createActorImageContainerProperties,
		arg.AIID,
		arg.IsContainer,
		arg.IsSurfaceContainer,
		arg.LiquidCapacity,
	)
}

const createActorImageFoodProperties = `-- name: CreateActorImageFoodProperties :execresult
INSERT INTO actor_images_food_properties (aiid, eats_into, sustenance) VALUES (?, ?, ?)
`

type CreateActorImageFoodPropertiesParams struct {
	AIID       int64
	EatsInto   int64
	Sustenance int32
}

func (q *Queries) CreateActorImageFoodProperties(ctx context.Context, arg CreateActorImageFoodPropertiesParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageFoodPropertiesStmt, createActorImageFoodProperties, arg.AIID, arg.EatsInto, arg.Sustenance)
}

const createActorImageFurnitureProperties = `-- name: CreateActorImageFurnitureProperties :execresult
INSERT INTO actor_images_furniture_properties (aiid, seating) VALUES (?, ?)
`

type CreateActorImageFurniturePropertiesParams struct {
	AIID    int64
	Seating int32
}

func (q *Queries) CreateActorImageFurnitureProperties(ctx context.Context, arg CreateActorImageFurniturePropertiesParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageFurniturePropertiesStmt, createActorImageFurnitureProperties, arg.AIID, arg.Seating)
}

const createActorImageHand = `-- name: CreateActorImageHand :execresult
INSERT INTO actor_images_hands (aiid, hand) VALUES (?, ?)
`

type CreateActorImageHandParams struct {
	AIID int64
	Hand int32
}

func (q *Queries) CreateActorImageHand(ctx context.Context, arg CreateActorImageHandParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageHandStmt, createActorImageHand, arg.AIID, arg.Hand)
}

const createActorImageKeyword = `-- name: CreateActorImageKeyword :execresult
INSERT INTO actor_images_keywords (keyword, aiid) VALUES (?, ?)
`

type CreateActorImageKeywordParams struct {
	Keyword string
	AIID    int64
}

func (q *Queries) CreateActorImageKeyword(ctx context.Context, arg CreateActorImageKeywordParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImageKeywordStmt, createActorImageKeyword, arg.Keyword, arg.AIID)
}

const createActorImagePrimaryHand = `-- name: CreateActorImagePrimaryHand :execresult
INSERT INTO actor_images_primary_hands (aiid, hand) VALUES (?, ?)
`

type CreateActorImagePrimaryHandParams struct {
	AIID int64
	Hand int32
}

func (q *Queries) CreateActorImagePrimaryHand(ctx context.Context, arg CreateActorImagePrimaryHandParams) (sql.Result, error) {
	return q.exec(ctx, q.createActorImagePrimaryHandStmt, createActorImagePrimaryHand, arg.AIID, arg.Hand)
}

const deleteActorImageCan = `-- name: DeleteActorImageCan :exec
DELETE FROM actor_images_can WHERE id = ?
`

func (q *Queries) DeleteActorImageCan(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteActorImageCanStmt, deleteActorImageCan, id)
	return err
}

const deleteActorImageCanBe = `-- name: DeleteActorImageCanBe :exec
DELETE FROM actor_images_can_be WHERE id = ?
`

func (q *Queries) DeleteActorImageCanBe(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteActorImageCanBeStmt, deleteActorImageCanBe, id)
	return err
}

const deleteActorImageContainerProperties = `-- name: DeleteActorImageContainerProperties :exec
DELETE FROM actor_images_container_properties WHERE id = ?
`

func (q *Queries) DeleteActorImageContainerProperties(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteActorImageContainerPropertiesStmt, deleteActorImageContainerProperties, id)
	return err
}

const deleteActorImageFoodProperties = `-- name: DeleteActorImageFoodProperties :exec
DELETE FROM actor_images_food_properties WHERE id = ?
`

func (q *Queries) DeleteActorImageFoodProperties(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteActorImageFoodPropertiesStmt, deleteActorImageFoodProperties, id)
	return err
}

const deleteActorImageFurnitureProperties = `-- name: DeleteActorImageFurnitureProperties :exec
DELETE FROM actor_images_furniture_properties WHERE id = ?
`

func (q *Queries) DeleteActorImageFurnitureProperties(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteActorImageFurniturePropertiesStmt, deleteActorImageFurnitureProperties, id)
	return err
}

const deleteActorImageHand = `-- name: DeleteActorImageHand :exec
DELETE FROM actor_images_hands WHERE id = ?
`

func (q *Queries) DeleteActorImageHand(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteActorImageHandStmt, deleteActorImageHand, id)
	return err
}

const deleteActorImagePrimaryHand = `-- name: DeleteActorImagePrimaryHand :exec
DELETE FROM actor_images_primary_hands WHERE id = ?
`

func (q *Queries) DeleteActorImagePrimaryHand(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteActorImagePrimaryHandStmt, deleteActorImagePrimaryHand, id)
	return err
}

const getActorImage = `-- name: GetActorImage :one
SELECT created_at, updated_at, description, short_description, name, gender, id, uniq FROM actor_images WHERE id = ?
`

func (q *Queries) GetActorImage(ctx context.Context, id int64) (ActorImage, error) {
	row := q.queryRow(ctx, q.getActorImageStmt, getActorImage, id)
	var i ActorImage
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.ShortDescription,
		&i.Name,
		&i.Gender,
		&i.ID,
		&i.Uniq,
	)
	return i, err
}

const getActorImageByName = `-- name: GetActorImageByName :one
SELECT created_at, updated_at, description, short_description, name, gender, id, uniq FROM actor_images WHERE name = ?
`

func (q *Queries) GetActorImageByName(ctx context.Context, name string) (ActorImage, error) {
	row := q.queryRow(ctx, q.getActorImageByNameStmt, getActorImageByName, name)
	var i ActorImage
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.ShortDescription,
		&i.Name,
		&i.Gender,
		&i.ID,
		&i.Uniq,
	)
	return i, err
}

const getActorImageContainerProperties = `-- name: GetActorImageContainerProperties :one
SELECT created_at, updated_at, aiid, id, liquid_capacity, is_container, is_surface_container FROM actor_images_container_properties WHERE aiid = ?
`

func (q *Queries) GetActorImageContainerProperties(ctx context.Context, aiid int64) (ActorImagesContainerProperty, error) {
	row := q.queryRow(ctx, q.getActorImageContainerPropertiesStmt, getActorImageContainerProperties, aiid)
	var i ActorImagesContainerProperty
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AIID,
		&i.ID,
		&i.LiquidCapacity,
		&i.IsContainer,
		&i.IsSurfaceContainer,
	)
	return i, err
}

const getActorImageFoodProperties = `-- name: GetActorImageFoodProperties :one
SELECT created_at, updated_at, eats_into, aiid, id, sustenance FROM actor_images_food_properties WHERE aiid = ?
`

func (q *Queries) GetActorImageFoodProperties(ctx context.Context, aiid int64) (ActorImagesFoodProperty, error) {
	row := q.queryRow(ctx, q.getActorImageFoodPropertiesStmt, getActorImageFoodProperties, aiid)
	var i ActorImagesFoodProperty
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EatsInto,
		&i.AIID,
		&i.ID,
		&i.Sustenance,
	)
	return i, err
}

const getActorImageFurnitureProperties = `-- name: GetActorImageFurnitureProperties :one
SELECT created_at, updated_at, eats_into, aiid, id, sustenance FROM actor_images_food_properties WHERE aiid = ?
`

func (q *Queries) GetActorImageFurnitureProperties(ctx context.Context, aiid int64) (ActorImagesFoodProperty, error) {
	row := q.queryRow(ctx, q.getActorImageFurniturePropertiesStmt, getActorImageFurnitureProperties, aiid)
	var i ActorImagesFoodProperty
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EatsInto,
		&i.AIID,
		&i.ID,
		&i.Sustenance,
	)
	return i, err
}

const listActorImageCan = `-- name: ListActorImageCan :many
SELECT created_at, updated_at, can, aiid, id FROM actor_images_can WHERE aiid = ?
`

func (q *Queries) ListActorImageCan(ctx context.Context, aiid int64) ([]ActorImagesCan, error) {
	rows, err := q.query(ctx, q.listActorImageCanStmt, listActorImageCan, aiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ActorImagesCan
	for rows.Next() {
		var i ActorImagesCan
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Can,
			&i.AIID,
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

const listActorImageCanBe = `-- name: ListActorImageCanBe :many
SELECT created_at, updated_at, can_be, aiid, id FROM actor_images_can_be WHERE aiid = ?
`

func (q *Queries) ListActorImageCanBe(ctx context.Context, aiid int64) ([]ActorImagesCanBe, error) {
	rows, err := q.query(ctx, q.listActorImageCanBeStmt, listActorImageCanBe, aiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ActorImagesCanBe
	for rows.Next() {
		var i ActorImagesCanBe
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CanBe,
			&i.AIID,
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

const listActorImageKeywords = `-- name: ListActorImageKeywords :many
SELECT created_at, updated_at, keyword, aiid, id FROM actor_images_keywords WHERE aiid = ?
`

func (q *Queries) ListActorImageKeywords(ctx context.Context, aiid int64) ([]ActorImagesKeyword, error) {
	rows, err := q.query(ctx, q.listActorImageKeywordsStmt, listActorImageKeywords, aiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ActorImagesKeyword
	for rows.Next() {
		var i ActorImagesKeyword
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Keyword,
			&i.AIID,
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

const listActorImages = `-- name: ListActorImages :many
SELECT created_at, updated_at, description, short_description, name, gender, id, uniq FROM actor_images
`

func (q *Queries) ListActorImages(ctx context.Context) ([]ActorImage, error) {
	rows, err := q.query(ctx, q.listActorImagesStmt, listActorImages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ActorImage
	for rows.Next() {
		var i ActorImage
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
			&i.ShortDescription,
			&i.Name,
			&i.Gender,
			&i.ID,
			&i.Uniq,
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

const listActorImagesHands = `-- name: ListActorImagesHands :many
SELECT created_at, updated_at, aiid, id, hand FROM actor_images_hands WHERE aiid = ?
`

func (q *Queries) ListActorImagesHands(ctx context.Context, aiid int64) ([]ActorImagesHand, error) {
	rows, err := q.query(ctx, q.listActorImagesHandsStmt, listActorImagesHands, aiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ActorImagesHand
	for rows.Next() {
		var i ActorImagesHand
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AIID,
			&i.ID,
			&i.Hand,
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

const listActorImagesPrimaryHands = `-- name: ListActorImagesPrimaryHands :many
SELECT created_at, updated_at, aiid, id, hand FROM actor_images_primary_hands WHERE aiid = ?
`

func (q *Queries) ListActorImagesPrimaryHands(ctx context.Context, aiid int64) ([]ActorImagesPrimaryHand, error) {
	rows, err := q.query(ctx, q.listActorImagesPrimaryHandsStmt, listActorImagesPrimaryHands, aiid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ActorImagesPrimaryHand
	for rows.Next() {
		var i ActorImagesPrimaryHand
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AIID,
			&i.ID,
			&i.Hand,
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
