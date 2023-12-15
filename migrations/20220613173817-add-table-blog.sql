
-- +migrate Up
CREATE TABLE IF NOT EXISTS Blogs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    photo VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE INDEX IF NOT EXISTS idx_title ON Blogs(title);
CREATE INDEX IF NOT EXISTS idx_blog_user ON Blogs(user_id);
-- +migrate Down
DROP TABLE IF EXISTS Blogs;
