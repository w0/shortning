-- name: NewUrl :one
INSERT INTO urls (created_at, updated_at, url)
VALUES (
    NOW(),
    NOW(),
    $1
)
RETURNING *;

-- name: GetUrl :one
SELECT * FROM urls WHERE id = $1;

-- name: AddClick :exec
UPDATE urls
SET clicks = clicks + 1,
    updated_at = NOW()
WHERE id = $1;


-- name: GetUrlsCreatedBefore :many
SELECT id FROM urls
WHERE created_at < now() - MAKE_INTERVAL(DAYS => $1);

-- name: GetUrlsUnderClickCount :many
SELECT * FROM urls
WHERE clicks <= $1;

-- name: DeleteUrl :exec
DELETE FROM urls
WHERE id = $1;
