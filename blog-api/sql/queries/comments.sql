-- name: CreateComment :one
INSERT INTO comments (
    user_id,
    post_id,
    parent_comment_id,
    content
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetCommentById :one
SELECT * FROM comments WHERE id = $1 LIMIT 1;

-- name: ListCommentsByPostIdOrderByCreatedAtDesc :many
SELECT * FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY created_at DESC;

-- name: ListCommentsByPostIdOrderByCreatedAtAsc :many
SELECT * FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY created_at;

-- name: ListCommentsByPostIdOrderByLikesDesc :many
SELECT * FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY likes_count DESC, created_at DESC;

-- name: ListCommentsByPostIdOrderByLikesAsc :many
SELECT * FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY likes_count, created_at DESC;

-- name: ListCommentsRepliesOrderByCreateAtDesc :many
SELECT * FROM comments WHERE parent_comment_id = $1 ORDER BY created_at DESC;

-- name: ListCommentsRepliesOrderByCreateAtAsc :many
SELECT * FROM comments WHERE parent_comment_id = $1 ORDER BY created_at;

-- name: ListCommentsRepliesOrderByLikesDesc :many
SELECT * FROM comments WHERE parent_comment_id = $1 ORDER BY likes_count DESC, created_at DESC;

-- name: ListCommentsRepliesOrderByLikesAsc :many
SELECT * FROM comments WHERE parent_comment_id = $1 ORDER BY likes_count, created_at DESC;

-- name: ListCommentsByUserIdOrderByCreateAtDesc :many
SELECT * FROM comments WHERE user_id = $1 ORDER BY created_at DESC;

-- name: ListCommentsByUserIdOrderByCreateAtAsc :many
SELECT * FROM comments WHERE user_id = $1 ORDER BY created_at;

-- name: ListCommentsByUserIdOrderByLikesDesc :many
SELECT * FROM comments WHERE user_id = $1 ORDER BY likes_count DESC, created_at DESC;

-- name: ListCommentsByUserIdOrderByLikesAsc :many
SELECT * FROM comments WHERE user_id = $1 ORDER BY likes_count, created_at DESC;
