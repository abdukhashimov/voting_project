-- name: GetAllDistricts :many
SELECT *
FROM district
WHERE district.title ilike '%' || @search::VARCHAR || '%'
    AND CASE
        WHEN @region_id::INTEGER = 0 THEN TRUE
        ELSE district.region_id = @region_id::INTEGER
    END
LIMIT @limit_ OFFSET @offset_;

-- name: CreateDistrict :one
INSERT INTO district(title, region_id)
VALUES ($1, $2) RETURNING id;

-- name: GetDistrictByRegionID :many
SELECT *
FROM district
WHERE region_id = $1;

-- name: GetDistrictByID :one
SELECT *
FROM district
WHERE id = $1;

-- name: UpdateDistrict :exec
UPDATE district
SET title = COALESCE(sqlc.narg('title'), title),
    region_id = COALESCE(sqlc.narg('region_id'), region_id)
WHERE id = sqlc.arg('id');

-- name: DeleteDistrictByID :exec
DELETE FROM district
WHERE id = $1;

-- name: GetAllDistrictsCount :one
SELECT COUNT(1)
FROM district
WHERE district.title ilike '%' || @search::VARCHAR || '%'
    AND CASE
        WHEN @region_id::INTEGER = 0 THEN TRUE
        ELSE district.region_id = @region_id::INTEGER
    END;
