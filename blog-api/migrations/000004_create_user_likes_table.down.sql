ALTER TABLE IF EXISTS userLikes DROP CONSTRAINT IF EXISTS check_userlikes_target_type;

DROP INDEX IF EXISTS idx_userlikes_user_targets;
DROP INDEX IF EXISTS idx_userlikes_target;
DROP TABLE IF EXISTS UserLikes;