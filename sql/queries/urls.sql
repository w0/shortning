-- name: NewUrl :one
INSERT INTO urls (created_at, updated_at, url)
VALUES (
    NOW(),
    NOW(),
    $1
)
RETURNING *;
