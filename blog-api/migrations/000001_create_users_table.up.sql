CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    user_uuid UUID NOT NULL UNIQUE,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_user_uuid ON users(user_uuid);

COMMENT ON TABLE users IS 'Stores user account information';
COMMENT ON COLUMN users.id IS 'Primary Key, unique identifier for the user';
COMMENT ON COLUMN Users.user_uuid IS 'External unique identifier for the user (UUID), if using internal integer PK.';
COMMENT ON COLUMN Users.username IS 'Unique username for the user.';
COMMENT ON COLUMN Users.email IS 'Unique email address for the user.';
COMMENT ON COLUMN Users.password_hash IS 'Hashed password for the user.';
COMMENT ON COLUMN Users.created_at IS 'Timestamp of when the user was created.';
COMMENT ON COLUMN Users.updated_at IS 'Timestamp of when the user was last updated.'