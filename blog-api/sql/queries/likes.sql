-- name: CreateLike :one
INSERT INTO userlikes (
    user_id,
    target_id,
    target_type
) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteLike :exec
DELETE FROM userlikes WHERE  user_id = $1 AND target_id = $2 AND target_type = $3;

-- name: GetLike :one
SELECT * FROM userlikes WHERE user_id = $1 AND target_id = $2 AND target_type = $3 LIMIT 1;

-- name: CountLikesForTarget :one
SELECT COUNT(*) FROM userlikes WHERE target_id = $1 AND target_type = $2;

-- name: ListUserLikesByType :many
SELECT * FROM userlikes WHERE user_id = $1 AND target_type = $2 ORDER BY created_at DESC;