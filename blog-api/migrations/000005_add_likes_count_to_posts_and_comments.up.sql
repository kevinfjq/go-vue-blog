ALTER TABLE posts
    ADD COLUMN IF NOT EXISTS likes_count INTEGER NOT NULL DEFAULT 0;

COMMENT ON COLUMN posts.likes_count IS 'Cached count of likes for the post.';

ALTER TABLE comments
    ADD COLUMN IF NOT EXISTS likes_count INTEGER NOT NULL DEFAULT 0;

COMMENT ON COLUMN Comments.likes_count IS 'Cached count of likes for the comment.';