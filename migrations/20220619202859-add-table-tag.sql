
-- +migrate Up
CREATE TABLE IF NOT EXISTS Tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tag VARCHAR(255) UNIQUE NOT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE INDEX IF NOT EXISTS idx_tags ON Tags(tag);

-- +migrate Down
DROP TABLE IF EXISTS Tags;