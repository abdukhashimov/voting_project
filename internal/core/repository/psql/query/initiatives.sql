-- name: CreateInitiative :one
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
        sqlc.arg(title),
        sqlc.arg(images),
        sqlc.arg(description),
        sqlc.arg(author),
        sqlc.arg(board_id),
        sqlc.arg(vote_count),
        sqlc.arg(requested_amount),
        sqlc.arg(granted_amount),
        sqlc.arg(region_id),
        sqlc.arg(district_id)
    ) RETURNING id;

-- name: GetInitiativeByID :one
SELECT *
FROM initiatives
WHERE id = sqlc.arg('id');

-- name: UpdateInitiativeByID :exec
UPDATE initiatives
SET title = COALESCE(sqlc.narg('title'), title),
    images = COALESCE(sqlc.narg('images'), images),
    description = COALESCE(sqlc.narg('description'), description),
    vote_count = COALESCE(sqlc.narg('vote_count'), vote_count),
    requested_amount = COALESCE(sqlc.narg('requested_amount'), requested_amount),
    granted_amount = COALESCE(sqlc.narg('granted_amount'), granted_amount),
    region_id = COALESCE(sqlc.narg('region_id'), region_id),
    district_id = COALESCE(sqlc.narg('district_id'), district_id),
    updated_at = now()
WHERE id = sqlc.arg('id');

-- name: GetAllInitiatives :many
SELECT *
FROM initiatives
WHERE initiatives.title ilike '%' || @search::VARCHAR || '%'
    AND CASE
        WHEN @region_id::INTEGER = 0 THEN TRUE
        ELSE initiatives.region_id = @region_id::INTEGER
    END
    AND CASE
        WHEN @district_id::INTEGER = 0 THEN TRUE
        ELSE initiatives.district_id = @district_id::INTEGER
    END
    AND CASE
        WHEN @granted_amount_start::INTEGER = 0 THEN TRUE
        ELSE initiatives.granted_amount_start >= @granted_amount_start::INTEGER
    END
    AND CASE
        WHEN @granted_amount_end::INTEGER = 0 THEN TRUE
        ELSE initiatives.granted_amount_end <= @granted_amount_end::INTEGER
    END
    AND CASE
        WHEN @vote_count_start::INTEGER = 0 THEN TRUE
        ELSE initiatives.vote_count_start >= @vote_count_start::INTEGER
    END
    AND CASE
        WHEN @vote_count_end::INTEGER = 0 THEN TRUE
        ELSE initiatives.vote_count_end <= @vote_count_end::INTEGER
    END
    AND CASE
        WHEN @created_at_start::TIMESTAMP IS NULL THEN TRUE
        ELSE initiatives.created_at >= @created_at_start::TIMESTAMP
    END
    AND CASE
        WHEN @created_at_end::TIMESTAMP IS NULL THEN TRUE
        ELSE initiatives.created_at <= @created_at_end::TIMESTAMP
    END
    AND CASE
        WHEN @board_id::VARCHAR = '' THEN TRUE
        ELSE initiatives.board_id <= @board_id::VARCHAR
    END
    AND CASE
        WHEN @author::VARCHAR = '' THEN TRUE
        ELSE initiatives.author <= @author::VARCHAR
    END
ORDER BY created_at,
    vote_count DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetAllIniativesCount :one
SELECT COUNT(1)
FROM initiatives
WHERE initiatives.title ilike '%' || @search::VARCHAR || '%'
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

-- name: SoftDeleteInitiative :exec
UPDATE initiatives
SET status = @status::INTEGER,
    updated_at = now()
WHERE id = sqlc.arg('id');