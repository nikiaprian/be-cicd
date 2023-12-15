
-- +migrate Up
CREATE TABLE IF NOT EXISTS BlogTags (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tag_id INTEGER NOT NULL,
    blog_id INTEGER NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

    FOREIGN KEY (tag_id) REFERENCES Tags (id)
    FOREIGN KEY (blog_id) REFERENCES Blogs (id)
);

CREATE INDEX IF NOT EXISTS idx_tag_blog_tag ON BlogTags(tag_id);
CREATE INDEX IF NOT EXISTS idx_blog_blog_tag ON BlogTags(blog_id);
-- +migrate Down
DROP TABLE IF EXISTS BlogTags;
