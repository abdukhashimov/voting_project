-- name: CreateDistrictBudget :exec
INSERT INTO district_budget (board_id, district_id, budget, updated_by)
VALUES (
        sqlc.arg(board_id),
        sqlc.arg(district_id),
        sqlc.arg(budget),
        sqlc.arg(updated_by)
    );

-- name: UpdateDistrictBudget :exec
UPDATE district_budget
SET budget = sqlc.arg(budget),
    updated_by = sqlc.arg(updated_by),
    updated_at = now()
WHERE board_id = sqlc.arg(board_id)
    AND district_id = sqlc.arg(district_id);

-- name: DeleteDistrictBoard :exec
DELETE FROM district_budget
WHERE board_id = sqlc.arg(board_id)
    AND district_id = sqlc.arg(district_id);

-- name: GetAllDistrictBudget :many
SELECT *
FROM district_budget
WHERE TRUE
    AND CASE
        WHEN @board_id::INTEGER = 0 THEN TRUE
        ELSE district_budget.board_id = @board_id::INTEGER
    END
    AND CASE
        WHEN @district_id::INTEGER = 0 THEN TRUE
        ELSE district_budget.district_id = @district_id::INTEGER
    END
ORDER BY created_at DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');