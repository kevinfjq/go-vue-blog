CREATE TABLE IF NOT EXISTS userLikes (
    user_id BIGINT NOT NULL,
    target_id BIGINT NOT NULL,
    target_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, target_id, target_type),

    CONSTRAINT fk_userLikes_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_userLikes_target ON userLikes(target_id, target_type);
CREATE INDEX IF NOT EXISTS idx_userLikes_user_targets ON userLikes(user_id, target_type);

ALTER TABLE UserLikes ADD CONSTRAINT check_userlikes_target_type CHECK (target_type IN ('post', 'comment'));

COMMENT ON TABLE UserLikes IS 'Stores likes given by users to targets (e.g., posts, comments).';
COMMENT ON COLUMN UserLikes.user_id IS 'Foreign key referencing the user (Users.id) who liked the item.';
COMMENT ON COLUMN UserLikes.target_id IS 'Identifier of the item that was liked (Posts.id or Comments.id).';
COMMENT ON COLUMN UserLikes.target_type IS 'Type of the item that was liked (e.g., ''post'', ''comment'').';
COMMENT ON COLUMN UserLikes.created_at IS 'Timestamp of when the like was created.';