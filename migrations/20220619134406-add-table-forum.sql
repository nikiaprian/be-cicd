
-- +migrate Up
CREATE TABLE IF NOT EXISTS Forums (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE INDEX IF NOT EXISTS idx_forum_title ON Forums(title);
CREATE INDEX IF NOT EXISTS idx_forum_user ON Forums(user_id);
-- +migrate Down
DROP TABLE IF EXISTS Forums;
