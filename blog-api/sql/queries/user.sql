-- name: CreateUser :one
INSERT INTO users (
                user_uuid,
                username,
                email,
                password_hash
) VALUES (
    $1, $2, $3, $4
 ) RETURNING *;

-- name: GetUserByUUID :one
SELECT * FROM users WHERE user_uuid = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: UpdateUserPasswordById :one
UPDATE users SET password_hash = $2, updated_at = NOW() WHERE id = $1  RETURNING *;

-- name: UpdateUserEmailById :one
UPDATE users SET email = $2, updated_at = NOW() WHERE id = $1 RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;

