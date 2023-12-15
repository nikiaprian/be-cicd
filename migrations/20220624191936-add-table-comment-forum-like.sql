
-- +migrate Up
CREATE TABLE IF NOT EXISTS ForumCommentLikes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    forum_comment_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,

    FOREIGN KEY (forum_comment_id) REFERENCES CommentForum (id)
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE INDEX IF NOT EXISTS idx_forum_comment_like ON ForumCommentLikes(forum_comment_id);
CREATE INDEX IF NOT EXISTS idx_user_forum_comment_like ON ForumCommentLikes(user_id);
-- +migrate Down
DROP TABLE IF EXISTS ForumCommentLikes;
