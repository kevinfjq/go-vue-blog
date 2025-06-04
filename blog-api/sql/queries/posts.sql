-- name: CreatePost :one
INSERT INTO posts (
    user_id,
    title,
    content
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetPostById :one
SELECT * FROM  posts WHERE id = $1;

-- name: ListPostsByUserIdOrderByCreatedAtDesc :many
SELECT * FROM posts WHERE user_id = $1 ORDER BY created_at;

-- name: ListPostsByUserIdOrderByCreatedAtAsc :many
SELECT * FROM posts WHERE user_id = $1 ORDER BY created_at;

-- name: ListPostsByUserIdOrderByLikesDesc :many
SELECT * FROM posts WHERE user_id = $1 ORDER BY likes_count DESC, created_at DESC;

-- name: ListPostsByUserIdOrderByLikesAsc :many
SELECT * FROM posts WHERE user_id = $1 ORDER BY likes_count, created_at DESC;

-- name: ListPostsOrderByCreatedAtDesc :many
SELECT * FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: ListPostsOrderByCreatedAtAsc :many
SELECT * FROM posts ORDER BY created_at  LIMIT $1 OFFSET $2;

-- name: ListPostsOrderByLikesAtDesc :many
SELECT * FROM posts ORDER BY likes_count DESC, created_at DESC LIMIT $1 OFFSET $2;

-- name: ListPostsOrderByLikesAtAsc :many
SELECT * FROM posts ORDER BY likes_count, created_at DESC LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts SET content = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;

-- name: IncrementPostLikesCount :one
UPDATE posts SET likes_count = likes_count + 1 WHERE id = $1 RETURNING *;

-- name: DecrementPostLikesCount :one
UPDATE posts SET likes_count = GREATEST(0, likes_count - 1) WHERE id = $1 RETURNING *;