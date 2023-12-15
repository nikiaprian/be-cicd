
-- +migrate Up
CREATE TABLE IF NOT EXISTS ForumTags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tag_id INTEGER NOT NULL,
    forum_id INTEGER NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

FOREIGN KEY (tag_id) REFERENCES Tags (id)
FOREIGN KEY (forum_id) REFERENCES Forums (id)
);

CREATE INDEX IF NOT EXISTS idx_tag_forum_tag ON ForumTags(tag_id);
CREATE INDEX IF NOT EXISTS idx_forum_forum_tag ON ForumTags(forum_id);

-- +migrate Down
DROP TABLE IF EXISTS ForumTags;
