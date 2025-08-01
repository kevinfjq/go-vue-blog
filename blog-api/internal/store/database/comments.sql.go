// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: comments.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
    user_id,
    post_id,
    parent_comment_id,
    content
) VALUES ($1, $2, $3, $4) RETURNING id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count
`

type CreateCommentParams struct {
	UserID          int64       `json:"user_id"`
	PostID          int64       `json:"post_id"`
	ParentCommentID pgtype.Int8 `json:"parent_comment_id"`
	Content         string      `json:"content"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment,
		arg.UserID,
		arg.PostID,
		arg.ParentCommentID,
		arg.Content,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.ParentCommentID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LikesCount,
	)
	return i, err
}

const getCommentById = `-- name: GetCommentById :one
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCommentById(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRow(ctx, getCommentById, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.ParentCommentID,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LikesCount,
	)
	return i, err
}

const listCommentsByPostIdOrderByCreatedAtAsc = `-- name: ListCommentsByPostIdOrderByCreatedAtAsc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY created_at
`

func (q *Queries) ListCommentsByPostIdOrderByCreatedAtAsc(ctx context.Context, postID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByPostIdOrderByCreatedAtAsc, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByPostIdOrderByCreatedAtDesc = `-- name: ListCommentsByPostIdOrderByCreatedAtDesc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY created_at DESC
`

func (q *Queries) ListCommentsByPostIdOrderByCreatedAtDesc(ctx context.Context, postID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByPostIdOrderByCreatedAtDesc, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByPostIdOrderByLikesAsc = `-- name: ListCommentsByPostIdOrderByLikesAsc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY likes_count, created_at DESC
`

func (q *Queries) ListCommentsByPostIdOrderByLikesAsc(ctx context.Context, postID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByPostIdOrderByLikesAsc, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByPostIdOrderByLikesDesc = `-- name: ListCommentsByPostIdOrderByLikesDesc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE post_id = $1 AND parent_comment_id IS NULL ORDER BY likes_count DESC, created_at DESC
`

func (q *Queries) ListCommentsByPostIdOrderByLikesDesc(ctx context.Context, postID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByPostIdOrderByLikesDesc, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByUserIdOrderByCreateAtAsc = `-- name: ListCommentsByUserIdOrderByCreateAtAsc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE user_id = $1 ORDER BY created_at
`

func (q *Queries) ListCommentsByUserIdOrderByCreateAtAsc(ctx context.Context, userID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByUserIdOrderByCreateAtAsc, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByUserIdOrderByCreateAtDesc = `-- name: ListCommentsByUserIdOrderByCreateAtDesc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE user_id = $1 ORDER BY created_at DESC
`

func (q *Queries) ListCommentsByUserIdOrderByCreateAtDesc(ctx context.Context, userID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByUserIdOrderByCreateAtDesc, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByUserIdOrderByLikesAsc = `-- name: ListCommentsByUserIdOrderByLikesAsc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE user_id = $1 ORDER BY likes_count, created_at DESC
`

func (q *Queries) ListCommentsByUserIdOrderByLikesAsc(ctx context.Context, userID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByUserIdOrderByLikesAsc, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsByUserIdOrderByLikesDesc = `-- name: ListCommentsByUserIdOrderByLikesDesc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE user_id = $1 ORDER BY likes_count DESC, created_at DESC
`

func (q *Queries) ListCommentsByUserIdOrderByLikesDesc(ctx context.Context, userID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByUserIdOrderByLikesDesc, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsRepliesOrderByCreateAtAsc = `-- name: ListCommentsRepliesOrderByCreateAtAsc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE parent_comment_id = $1 ORDER BY created_at
`

func (q *Queries) ListCommentsRepliesOrderByCreateAtAsc(ctx context.Context, parentCommentID pgtype.Int8) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsRepliesOrderByCreateAtAsc, parentCommentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsRepliesOrderByCreateAtDesc = `-- name: ListCommentsRepliesOrderByCreateAtDesc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE parent_comment_id = $1 ORDER BY created_at DESC
`

func (q *Queries) ListCommentsRepliesOrderByCreateAtDesc(ctx context.Context, parentCommentID pgtype.Int8) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsRepliesOrderByCreateAtDesc, parentCommentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsRepliesOrderByLikesAsc = `-- name: ListCommentsRepliesOrderByLikesAsc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE parent_comment_id = $1 ORDER BY likes_count, created_at DESC
`

func (q *Queries) ListCommentsRepliesOrderByLikesAsc(ctx context.Context, parentCommentID pgtype.Int8) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsRepliesOrderByLikesAsc, parentCommentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCommentsRepliesOrderByLikesDesc = `-- name: ListCommentsRepliesOrderByLikesDesc :many
SELECT id, post_id, user_id, parent_comment_id, content, created_at, updated_at, likes_count FROM comments WHERE parent_comment_id = $1 ORDER BY likes_count DESC, created_at DESC
`

func (q *Queries) ListCommentsRepliesOrderByLikesDesc(ctx context.Context, parentCommentID pgtype.Int8) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsRepliesOrderByLikesDesc, parentCommentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.ParentCommentID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LikesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
