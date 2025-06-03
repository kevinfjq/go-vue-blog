CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL PRIMARY KEY,
    post_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    parent_comment_id BIGINT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_comments_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,

    CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT fk_comments_parent_comment FOREIGN KEY (parent_comment_id) REFERENCES comments(id) ON DELETE CASCADE
);


CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);
CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
CREATE INDEX IF NOT EXISTS idx_comments_parent_comment_id ON comments(parent_comment_id);

COMMENT ON TABLE Comments IS 'Stores comments made by users on posts or other comments.';
COMMENT ON COLUMN Comments.id IS 'Internal primary key (auto-incrementing big integer). Will be exposed externally via HashID.';
COMMENT ON COLUMN Comments.post_id IS 'Foreign key referencing the post (Posts.id) this comment belongs to.';
COMMENT ON COLUMN Comments.user_id IS 'Foreign key referencing the user (Users.id) who wrote the comment.';
COMMENT ON COLUMN Comments.parent_comment_id IS 'Foreign key referencing another comment (Comments.id) if this is a reply. NULL for top-level comments.';
COMMENT ON COLUMN Comments.content IS 'The text content of the comment.';
COMMENT ON COLUMN Comments.created_at IS 'Timestamp of when the comment was created.';
COMMENT ON COLUMN Comments.updated_at IS 'Timestamp of when the comment was last updated.';