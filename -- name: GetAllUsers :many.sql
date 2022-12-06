-- name: GetAllUsers :many
select
    *
from
    users
WHERE
    users.fullname ilike '%' || @search :: varchar || '%'
    AND CASE
        WHEN @region_id :: varchar = '' THEN true
        ELSE users.region_id = @region_id :: varchar
    END
ORDER BY
    created_at DESC OFFSET @offset :: integer
LIMIT
    @limit :: integer;
