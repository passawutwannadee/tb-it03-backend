-- name: IT03List :many
SELECT 
    i.id,
    i.name,
    i.reason,
    i.status_id,
    s.status
FROM it03 i
LEFT JOIN it03_statuses s ON i.status_id = s.id
WHERE deleted_at is NULL;

-- name: IT03Count :one
SELECT COUNT(*) FROM it03 WHERE deleted_at is NULL;

-- name: IT03Update :many
WITH it03 AS (
    UPDATE it03
    SET
        name = COALESCE(sqlc.narg('name'), name),
        reason = COALESCE(sqlc.narg('reason'), reason),
        status_id = COALESCE(sqlc.narg('status_id'), status_id)
    WHERE id = ANY($1::int4[])
      AND deleted_at IS NULL
      AND it03.status_id = COALESCE(sqlc.narg('with_status_id'), it03.status_id)
    RETURNING *
)
SELECT 
    u.id,
    u.name,
    u.reason,
    u.status_id,
    s.status
FROM it03 u
LEFT JOIN it03_statuses s ON u.status_id = s.id;
