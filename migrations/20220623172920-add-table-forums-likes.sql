
-- +migrate Up
CREATE TABLE IF NOT EXISTS ForumsLikes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    forum_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

    FOREIGN KEY (forum_id) REFERENCES Forums (id)
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE INDEX IF NOT EXISTS idx_forum_forum_like ON ForumsLikes(forum_id);
CREATE INDEX IF NOT EXISTS idx_user_forum_like ON ForumsLikes(user_id);
-- +migrate Down
DROP TABLE IF EXISTS ForumsLikes;
