
-- +migrate Up
CREATE TABLE IF NOT EXISTS BlogCommentLikes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    blog_comment_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

    FOREIGN KEY (blog_comment_id) REFERENCES CommentBlog (id)
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE INDEX IF NOT EXISTS idx_blog_comment_like ON BlogCommentLikes(blog_comment_id);
CREATE INDEX IF NOT EXISTS idx_user_blog_comment_like ON BlogCommentLikes(user_id);
-- +migrate Down
DROP TABLE IF EXISTS BlogCommentLikes;
