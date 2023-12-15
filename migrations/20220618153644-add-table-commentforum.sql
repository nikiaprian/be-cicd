
-- +migrate Up
CREATE TABLE IF NOT EXISTS CommentForum (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    forum_id INTEGER NOT NULL,
    comment TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    
    FOREIGN KEY (user_id) REFERENCES Users (id),
    FOREIGN KEY (forum_id) REFERENCES Forums (id)
);

CREATE INDEX IF NOT EXISTS idx_commentforum_user ON CommentForum(user_id);
CREATE INDEX IF NOT EXISTS idx_commentforum_forum ON CommentForum(forum_id);

-- +migrate Down
DROP TABLE IF EXISTS CommentForum;