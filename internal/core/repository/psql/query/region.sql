-- name: GetAllRegions :many
SELECT *
FROM region
WHERE region.title ilike '%' || @search::VARCHAR || '%'
LIMIT @limit_ OFFSET @offset_;

-- name: GetRegionByID :one
SELECT *
FROM region
WHERE id = $1;

-- name: UpdateRegion :exec
UPDATE region
SET title = $1
WHERE id = $2;

-- name: DeleteRegion :exec
DELETE FROM region
WHERE id = $1;

-- name: CreateRegion :one
INSERT INTO region (title)
VALUES ($1) RETURNING id;

-- name: GetAllRegionsCount :one
SELECT COUNT(1)
FROM region
WHERE region.title ilike '%' || @search::VARCHAR || '%';
