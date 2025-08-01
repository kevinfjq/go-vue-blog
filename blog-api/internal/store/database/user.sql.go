// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: user.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
                user_uuid,
                username,
                email,
                password_hash
) VALUES (
    $1, $2, $3, $4
 ) RETURNING id, user_uuid, username, email, password_hash, created_at, updated_at
`

type CreateUserParams struct {
	UserUuid     pgtype.UUID `json:"user_uuid"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserUuid,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserById = `-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUserById(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUserById, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, user_uuid, username, email, password_hash, created_at, updated_at FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, user_uuid, username, email, password_hash, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUUID = `-- name: GetUserByUUID :one
SELECT id, user_uuid, username, email, password_hash, created_at, updated_at FROM users WHERE user_uuid = $1 LIMIT 1
`

func (q *Queries) GetUserByUUID(ctx context.Context, userUuid pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUUID, userUuid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, user_uuid, username, email, password_hash, created_at, updated_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserEmailById = `-- name: UpdateUserEmailById :one
UPDATE users SET email = $2, updated_at = NOW() WHERE id = $1 RETURNING id, user_uuid, username, email, password_hash, created_at, updated_at
`

type UpdateUserEmailByIdParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUserEmailById(ctx context.Context, arg UpdateUserEmailByIdParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserEmailById, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPasswordById = `-- name: UpdateUserPasswordById :one
UPDATE users SET password_hash = $2, updated_at = NOW() WHERE id = $1  RETURNING id, user_uuid, username, email, password_hash, created_at, updated_at
`

type UpdateUserPasswordByIdParams struct {
	ID           int64  `json:"id"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) UpdateUserPasswordById(ctx context.Context, arg UpdateUserPasswordByIdParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserPasswordById, arg.ID, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
