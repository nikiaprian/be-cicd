
-- +migrate Up
CREATE TABLE IF NOT EXISTS CommentBlog (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  blog_id INTEGER NOT NULL,
  comment TEXT NOT NULL,
  created_at DATETIME,
  updated_at DATETIME,

  FOREIGN KEY (user_id) REFERENCES Users (id),
  FOREIGN KEY (blog_id) REFERENCES Blogs (id)
);

CREATE INDEX IF NOT EXISTS idx_commentblog_user ON CommentBlog(user_id);
CREATE INDEX IF NOT EXISTS idx_commentblog_blog ON CommentBlog(blog_id);
-- +migrate Down
DROP TABLE IF EXISTS CommentBlog;
