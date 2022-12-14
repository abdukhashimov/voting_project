// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: initiatives.sql

package sqlc

import (
	"context"
	"time"

	null "gopkg.in/guregu/null.v4"
)

const createInitiative = `-- name: CreateInitiative :one
INSERT INTO initiatives (
        title,
        images,
        description,
        author,
        board_id,
        vote_count,
        requested_amount,
        granted_amount,
        region_id,
        district_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10
    ) RETURNING id
`

type CreateInitiativeParams struct {
	Title           string      `json:"title"`
	Images          []string    `json:"images"`
	Description     null.String `json:"description"`
	Author          string      `json:"author"`
	BoardID         string      `json:"board_id"`
	VoteCount       null.Int    `json:"vote_count"`
	RequestedAmount null.Int    `json:"requested_amount"`
	GrantedAmount   null.Int    `json:"granted_amount"`
	RegionID        int32       `json:"region_id"`
	DistrictID      int32       `json:"district_id"`
}

func (q *Queries) CreateInitiative(ctx context.Context, arg CreateInitiativeParams) (string, error) {
	row := q.db.QueryRow(ctx, createInitiative,
		arg.Title,
		arg.Images,
		arg.Description,
		arg.Author,
		arg.BoardID,
		arg.VoteCount,
		arg.RequestedAmount,
		arg.GrantedAmount,
		arg.RegionID,
		arg.DistrictID,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getAllIniativesCount = `-- name: GetAllIniativesCount :one
SELECT COUNT(1)
FROM initiatives
WHERE initiatives.title ilike '%' || $1::VARCHAR || '%'
    AND CASE
        WHEN $2::INTEGER = 0 THEN TRUE
        ELSE users.region_id = $2::INTEGER
    END
    AND CASE
        WHEN $3::INTEGER = 0 THEN TRUE
        ELSE users.district_id = $3::INTEGER
    END
    AND CASE
        WHEN $4::INTEGER = 0 THEN TRUE
        ELSE users.status = $4::INTEGER
    END
`

type GetAllIniativesCountParams struct {
	Search     string `json:"search"`
	RegionID   int32  `json:"region_id"`
	DistrictID int32  `json:"district_id"`
	Status     int32  `json:"status"`
}

func (q *Queries) GetAllIniativesCount(ctx context.Context, arg GetAllIniativesCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, getAllIniativesCount,
		arg.Search,
		arg.RegionID,
		arg.DistrictID,
		arg.Status,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAllInitiatives = `-- name: GetAllInitiatives :many
SELECT id, title, images, description, author, board_id, vote_count, status, requested_amount, granted_amount, region_id, district_id, quarter_id, created_at, updated_at
FROM initiatives
WHERE initiatives.title ilike '%' || $1::VARCHAR || '%'
    AND CASE
        WHEN $2::INTEGER = 0 THEN TRUE
        ELSE initiatives.region_id = $2::INTEGER
    END
    AND CASE
        WHEN $3::INTEGER = 0 THEN TRUE
        ELSE initiatives.district_id = $3::INTEGER
    END
    AND CASE
        WHEN $4::INTEGER = 0 THEN TRUE
        ELSE initiatives.granted_amount_start >= $4::INTEGER
    END
    AND CASE
        WHEN $5::INTEGER = 0 THEN TRUE
        ELSE initiatives.granted_amount_end <= $5::INTEGER
    END
    AND CASE
        WHEN $6::INTEGER = 0 THEN TRUE
        ELSE initiatives.vote_count_start >= $6::INTEGER
    END
    AND CASE
        WHEN $7::INTEGER = 0 THEN TRUE
        ELSE initiatives.vote_count_end <= $7::INTEGER
    END
    AND CASE
        WHEN $8::TIMESTAMP IS NULL THEN TRUE
        ELSE initiatives.created_at >= $8::TIMESTAMP
    END
    AND CASE
        WHEN $9::TIMESTAMP IS NULL THEN TRUE
        ELSE initiatives.created_at <= $9::TIMESTAMP
    END
    AND CASE
        WHEN $10::VARCHAR = '' THEN TRUE
        ELSE initiatives.board_id <= $10::VARCHAR
    END
    AND CASE
        WHEN $11::VARCHAR = '' THEN TRUE
        ELSE initiatives.author <= $11::VARCHAR
    END
ORDER BY created_at,
    vote_count DESC OFFSET $12
LIMIT $13
`

type GetAllInitiativesParams struct {
	Search             string    `json:"search"`
	RegionID           int32     `json:"region_id"`
	DistrictID         int32     `json:"district_id"`
	GrantedAmountStart int32     `json:"granted_amount_start"`
	GrantedAmountEnd   int32     `json:"granted_amount_end"`
	VoteCountStart     int32     `json:"vote_count_start"`
	VoteCountEnd       int32     `json:"vote_count_end"`
	CreatedAtStart     time.Time `json:"created_at_start"`
	CreatedAtEnd       time.Time `json:"created_at_end"`
	BoardID            string    `json:"board_id"`
	Author             string    `json:"author"`
	Offset             int32     `json:"offset"`
	Limit              int32     `json:"limit"`
}

func (q *Queries) GetAllInitiatives(ctx context.Context, arg GetAllInitiativesParams) ([]Initiative, error) {
	rows, err := q.db.Query(ctx, getAllInitiatives,
		arg.Search,
		arg.RegionID,
		arg.DistrictID,
		arg.GrantedAmountStart,
		arg.GrantedAmountEnd,
		arg.VoteCountStart,
		arg.VoteCountEnd,
		arg.CreatedAtStart,
		arg.CreatedAtEnd,
		arg.BoardID,
		arg.Author,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Initiative
	for rows.Next() {
		var i Initiative
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Images,
			&i.Description,
			&i.Author,
			&i.BoardID,
			&i.VoteCount,
			&i.Status,
			&i.RequestedAmount,
			&i.GrantedAmount,
			&i.RegionID,
			&i.DistrictID,
			&i.QuarterID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInitiativeByID = `-- name: GetInitiativeByID :one
SELECT id, title, images, description, author, board_id, vote_count, status, requested_amount, granted_amount, region_id, district_id, quarter_id, created_at, updated_at
FROM initiatives
WHERE id = $1
`

func (q *Queries) GetInitiativeByID(ctx context.Context, id string) (Initiative, error) {
	row := q.db.QueryRow(ctx, getInitiativeByID, id)
	var i Initiative
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Images,
		&i.Description,
		&i.Author,
		&i.BoardID,
		&i.VoteCount,
		&i.Status,
		&i.RequestedAmount,
		&i.GrantedAmount,
		&i.RegionID,
		&i.DistrictID,
		&i.QuarterID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const softDeleteInitiative = `-- name: SoftDeleteInitiative :exec
UPDATE initiatives
SET status = $1::INTEGER,
    updated_at = now()
WHERE id = $2
`

type SoftDeleteInitiativeParams struct {
	Status int32  `json:"status"`
	ID     string `json:"id"`
}

func (q *Queries) SoftDeleteInitiative(ctx context.Context, arg SoftDeleteInitiativeParams) error {
	_, err := q.db.Exec(ctx, softDeleteInitiative, arg.Status, arg.ID)
	return err
}

const updateInitiativeByID = `-- name: UpdateInitiativeByID :exec
UPDATE initiatives
SET title = COALESCE($1, title),
    images = COALESCE($2, images),
    description = COALESCE($3, description),
    vote_count = COALESCE($4, vote_count),
    requested_amount = COALESCE($5, requested_amount),
    granted_amount = COALESCE($6, granted_amount),
    region_id = COALESCE($7, region_id),
    district_id = COALESCE($8, district_id),
    updated_at = now()
WHERE id = $9
`

type UpdateInitiativeByIDParams struct {
	Title           null.String `json:"title"`
	Images          []string    `json:"images"`
	Description     null.String `json:"description"`
	VoteCount       null.Int    `json:"vote_count"`
	RequestedAmount null.Int    `json:"requested_amount"`
	GrantedAmount   null.Int    `json:"granted_amount"`
	RegionID        null.Int    `json:"region_id"`
	DistrictID      null.Int    `json:"district_id"`
	ID              string      `json:"id"`
}

func (q *Queries) UpdateInitiativeByID(ctx context.Context, arg UpdateInitiativeByIDParams) error {
	_, err := q.db.Exec(ctx, updateInitiativeByID,
		arg.Title,
		arg.Images,
		arg.Description,
		arg.VoteCount,
		arg.RequestedAmount,
		arg.GrantedAmount,
		arg.RegionID,
		arg.DistrictID,
		arg.ID,
	)
	return err
}
