-- name: GetAllUsers :many
SELECT *
FROM users
WHERE users.fullname ilike '%' || @search::VARCHAR || '%'
    AND CASE
        WHEN @region_id::INTEGER = 0 THEN TRUE
        ELSE users.region_id = @region_id::INTEGER
    END
    AND CASE
        WHEN @district_id::INTEGER = 0 THEN TRUE
        ELSE users.district_id = @district_id::INTEGER
    END
    AND CASE
        WHEN @status::INTEGER = 0 THEN TRUE
        ELSE users.status = @status::INTEGER
    END
ORDER BY created_at DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetAllUsersCount :one
SELECT COUNT(1)
FROM users
WHERE users.fullname ilike '%' || @search::VARCHAR || '%'
    AND CASE
        WHEN @region_id::INTEGER = 0 THEN TRUE
        ELSE users.region_id = @region_id::INTEGER
    END
    AND CASE
        WHEN @district_id::INTEGER = 0 THEN TRUE
        ELSE users.district_id = @district_id::INTEGER
    END
    AND CASE
        WHEN @status::INTEGER = 0 THEN TRUE
        ELSE users.status = @status::INTEGER
    END;

-- name: CreateUser :one
INSERT INTO users (fullname, phone_number, region_id, district_id)
VALUES ($1, $2, $3, $4) RETURNING id;

-- name: HardDeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET fullname = COALESCE(sqlc.narg('fullname'), fullname),
    phone_number = COALESCE(sqlc.narg('phone_number'), phone_number),
    region_id = COALESCE(sqlc.narg('region_id'), region_id),
    district_id = COALESCE(sqlc.narg('district_id'), district_id),
    username = COALESCE(sqlc.narg('username'), username),
    password = COALESCE(sqlc.narg('password'), password),
    role = COALESCE(sqlc.narg('role'), role),
    status = COALESCE(sqlc.narg('status'), status),
    updated_at = now()
WHERE id = sqlc.arg('id');

-- name: GetSingleUserByID :one
SELECT *
FROM users
WHERE id = sqlc.arg('id');

-- name: SoftDeleteUser :exec
UPDATE users
SET status = @status::INTEGER,
    updated_at = now()
WHERE id = sqlc.arg('id');