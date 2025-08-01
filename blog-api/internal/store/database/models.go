// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Stores comments made by users on posts or other comments.
type Comment struct {
	// Internal primary key (auto-incrementing big integer). Will be exposed externally via HashID.
	ID int64 `json:"id"`
	// Foreign key referencing the post (Posts.id) this comment belongs to.
	PostID int64 `json:"post_id"`
	// Foreign key referencing the user (Users.id) who wrote the comment.
	UserID int64 `json:"user_id"`
	// Foreign key referencing another comment (Comments.id) if this is a reply. NULL for top-level comments.
	ParentCommentID pgtype.Int8 `json:"parent_comment_id"`
	// The text content of the comment.
	Content string `json:"content"`
	// Timestamp of when the comment was created.
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	// Timestamp of when the comment was last updated.
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	// Cached count of likes for the comment.
	LikesCount int32 `json:"likes_count"`
}

// Stores blog posts or articles.
type Post struct {
	// Internal primary key (auto-incrementing big integer). Will be exposed externally via HashID.
	ID int64 `json:"id"`
	// Foreign key referencing the author (Users.id) of the post.
	UserID int64 `json:"user_id"`
	// Title of the post.
	Title string `json:"title"`
	// Main content of the post.
	Content pgtype.Text `json:"content"`
	// Timestamp of when the post was created.
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	// Timestamp of when the post was last updated.
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	// Cached count of likes for the post.
	LikesCount int32 `json:"likes_count"`
}

// Stores user account information
type User struct {
	// Primary Key, unique identifier for the user
	ID int64 `json:"id"`
	// External unique identifier for the user (UUID), if using internal integer PK.
	UserUuid pgtype.UUID `json:"user_uuid"`
	// Unique username for the user.
	Username string `json:"username"`
	// Unique email address for the user.
	Email string `json:"email"`
	// Hashed password for the user.
	PasswordHash string `json:"password_hash"`
	// Timestamp of when the user was created.
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	// Timestamp of when the user was last updated.
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

// Stores likes given by users to targets (e.g., posts, comments).
type Userlike struct {
	// Foreign key referencing the user (Users.id) who liked the item.
	UserID int64 `json:"user_id"`
	// Identifier of the item that was liked (Posts.id or Comments.id).
	TargetID int64 `json:"target_id"`
	// Type of the item that was liked (e.g., 'post', 'comment').
	TargetType string `json:"target_type"`
	// Timestamp of when the like was created.
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
