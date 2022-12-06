-- name: CreateBoard :one
INSERT INTO board(
        title,
        icon,
        total_amount,
        accept_start_date,
        accept_end_date,
        review_start_date,
        review_end_date,
        voting_start_date,
        voting_end_date
    )
VALUES (
        sqlc.arg('title'),
        sqlc.arg('icon'),
        sqlc.arg('total_amount'),
        sqlc.arg(accept_start_date),
        sqlc.arg(accept_end_date),
        sqlc.arg(review_start_date),
        sqlc.arg(review_end_date),
        sqlc.arg(voting_start_date),
        sqlc.arg(voting_end_date)
    ) RETURNING id;

-- name: GetAllBoards :many
SELECT *
FROM board
WHERE board.title ilike '%' || @search::VARCHAR || '%'
ORDER BY created_at DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');

-- name: GetAllBoardsCount :one
SELECT COUNT(1)
FROM board
WHERE board.title ilike '%' || @search::VARCHAR || '%';

-- name: GetBoardByID :one
SELECT *
FROM board
WHERE id = $1;

-- name: UpdateBoard :exec
UPDATE board
SET title = COALESCE(sqlc.narg('title'), title),
    icon = COALESCE(sqlc.narg('icon'), icon),
    total_amount = COALESCE(sqlc.narg('total_amount'), total_amount),
    accept_start_date = COALESCE(
        sqlc.narg('accept_start_date'),
        accept_start_date
    ),
    accept_end_date = COALESCE(sqlc.narg('accept_end_date'), accept_end_date),
    review_start_date = COALESCE(
        sqlc.narg('review_start_date'),
        review_start_date
    ),
    review_end_date = COALESCE(sqlc.narg('review_end_date'), review_end_date),
    voting_start_date = COALESCE(
        sqlc.narg('voting_start_date'),
        voting_start_date
    ),
    voting_end_date = COALESCE(sqlc.narg('voting_end_date'), voting_end_date),
    updated_at = now()
WHERE id = sqlc.arg('id');

-- name: HardDeleteBoard :exec
DELETE FROM board
WHERE id = $1;

-- name: SoftDeleteBoard :exec
UPDATE board
SET status = @status::INTEGER,
    updated_at = now()
WHERE id = sqlc.arg('id');