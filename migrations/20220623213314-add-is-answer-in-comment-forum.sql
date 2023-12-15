
-- +migrate Up
ALTER TABLE CommentForum
ADD is_answer boolean NOT NULL DEFAULT false;
-- +migrate Down
ALTER TABLE CommentForum
DROP CommentForum is_answer;