// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: urls.sql

package database

import (
	"context"
)

const addClick = `-- name: AddClick :exec
UPDATE urls
SET clicks = clicks + 1, updated_at = NOW()
WHERE id = $1
`

func (q *Queries) AddClick(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, addClick, id)
	return err
}

const getUrl = `-- name: GetUrl :one
SELECT id, created_at, updated_at, url, clicks FROM urls WHERE id = $1
`

func (q *Queries) GetUrl(ctx context.Context, id int32) (Url, error) {
	row := q.db.QueryRow(ctx, getUrl, id)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Clicks,
	)
	return i, err
}

const newUrl = `-- name: NewUrl :one
INSERT INTO urls (created_at, updated_at, url)
VALUES (
    NOW(),
    NOW(),
    $1
)
RETURNING id, created_at, updated_at, url, clicks
`

func (q *Queries) NewUrl(ctx context.Context, url string) (Url, error) {
	row := q.db.QueryRow(ctx, newUrl, url)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Clicks,
	)
	return i, err
}
