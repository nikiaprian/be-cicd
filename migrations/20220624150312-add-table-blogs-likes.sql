
-- +migrate Up
CREATE TABLE IF NOT EXISTS BlogsLikes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    blog_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

    FOREIGN KEY (blog_id) REFERENCES Blogs (id)
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE INDEX IF NOT EXISTS idx_blog_blog_like ON BlogsLikes(blog_id);
CREATE INDEX IF NOT EXISTS idx_user_blog_like ON BlogsLikes(user_id);
-- +migrate Down
DROP TABLE IF EXISTS BlogsLikes;
