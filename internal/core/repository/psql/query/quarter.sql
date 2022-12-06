-- name: GetAllQuarters :many
SELECT *
FROM quarter
WHERE quarter.title ilike '%' || @search::VARCHAR || '%'
    AND CASE
        WHEN @district_id::INTEGER = 0 THEN TRUE
        ELSE quarter.district_id = @district_id::INTEGER
    END
LIMIT @limit_ OFFSET @offset_;

-- name: CreateQuarter :one
INSERT INTO quarter(title, district_id)
VALUES ($1, $2) RETURNING id;

-- name: GetQuarterByID :one
SELECT *
FROM quarter
WHERE id = $1;

-- name: GetQuarterByDistrictID :many
SELECT *
FROM quarter
WHERE district_id = $1;

-- name: UpdateQuarter :exec
UPDATE quarter
SET title = COALESCE(sqlc.narg('title'), title),
    district_id = COALESCE(sqlc.narg('district_id'), disrtict_id)
WHERE id = sqlc.arg('id');

-- name: DeleteQuearterByID :exec
DELETE FROM quarter
WHERE id = $1;


-- name: GetAllQuartersCount :one
SELECT COUNT(1)
FROM quarter
WHERE quarter.title ilike '%' || @search::VARCHAR || '%'
    AND CASE
        WHEN @district_id::INTEGER = 0 THEN TRUE
        ELSE quarter.district_id = @district_id::INTEGER
    END;

