-- name: GetStatus :one
SELECT *
FROM status
WHERE id = $1;

-- name: CreateStatus :one
INSERT INTO status (title)
VALUES ($1) RETURNING id;

-- name: UpdateStatus :exec
UPDATE status
SET title = $1
WHERE id = $2;

-- name: DeleteStatus :exec
DELETE FROM status
WHERE id = $1;