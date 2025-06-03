CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_posts_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);

COMMENT ON TABLE Posts IS 'Stores blog posts or articles.';
COMMENT ON COLUMN Posts.id IS 'Internal primary key (auto-incrementing big integer). Will be exposed externally via HashID.';
COMMENT ON COLUMN Posts.user_id IS 'Foreign key referencing the author (Users.id) of the post.';
COMMENT ON COLUMN Posts.title IS 'Title of the post.';
COMMENT ON COLUMN Posts.content IS 'Main content of the post.';
COMMENT ON COLUMN Posts.created_at IS 'Timestamp of when the post was created.';
COMMENT ON COLUMN Posts.updated_at IS 'Timestamp of when the post was last updated.';