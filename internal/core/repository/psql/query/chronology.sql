-- name: CreateChronology :one
INSERT INTO chronology (title)
VALUES (sqlc.arg('title')) RETURNING id;

-- name: UpdateChronology :exec
UPDATE chronology
SET title = sqlc.arg('title'),
    updated_at = now()
WHERE id = sqlc.arg('id');

-- name: GetAllChronology :many
SELECT *
FROM chronology
WHERE chronology.title ilike '%' || @search::VARCHAR || '%'
ORDER BY created_at DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: DeleteChronology :exec
DELETE FROM chronology
WHERE id = sqlc.arg('id');