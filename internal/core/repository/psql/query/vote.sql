-- name: CreateVote :exec
INSERT INTO vote (phone_number, initiative_id, board_id)
VALUES (
        sqlc.arg(phone_number),
        sqlc.arg(initiative_id),
        sqlc.arg(board_id)
    );

-- name: DeleteVote :exec
DELETE FROM vote
WHERE phone_number = sqlc.arg('phone_number')
    AND board_id = sqlc.arg('board_id');

-- name: GetAllVotes :many
SELECT *
FROM vote
WHERE board_id = sqlc.arg('board_id')
ORDER BY created_at DESC OFFSET sqlc.arg('offset')
LIMIT sqlc.arg('limit');